# 동시성 패턴

![대표사진](../logo.png)

> 책 "Go In Action"을 공부하면서 정리한 문서입니다.

책에서는 동시성을 다루는 세 가지 패턴에 대해 다루고 있다.

1. 프로그램 생명 주기 관리하기
2. 재사용 가능한 리소스 풀 관리하기
3. 작업 처리 고루틴 풀 생성하기

각 패턴마다 그 예제를 살펴보자.


## 프로그램 생명 주기 관리하기 Runner

지금부터, `runner` 패키지를 만들어볼 것이다. 이 패키지는 실행 시간을 관찰하고 프로그램이 오래 실행되면, 프로그램을 종료시킨다. 백그라운드 작업 프로세스를 예약 실행 및 관리할 때 사용하면 좋은 패턴이다. (이런 프로그램을 `cron job`이라고 한다.)

먼저 실페 태스크들을 관리하는 `Runner`에 대한 코드를 만들 것이다.

```go
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner is ..
type Runner struct {
	// OS로부터 전달되는 인터럽트를 수신하기 위한 채널
	interrupt chan os.Signal
	// 처리가 종료되었음을 알리는 채널
	complete chan error
	// 지정된 시간이 초과했음을 알리는 채널
	timeout <-chan time.Time
	// 태스크를 위한 슬라이스
	tasks []func(int)
}

var ErrTimeout = errors.New("시간을 초과했습니다.")
var ErrInterrupt = errors.New("OS 인터럽트 신호를 수신했습니다.")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.goInterrupt() {
			return ErrInterrupt
		}

		task(id)
	}

	return nil
}

func (r *Runner) goInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
```

먼저 실제적으로 태스크를 관리할 수 있는 멤버 필드를 갖춘 `Runner` 구조체이다. 

```go
type Runner struct {
	// OS로부터 전달되는 인터럽트를 수신하기 위한 채널
	interrupt chan os.Signal
	// 처리가 종료되었음을 알리는 채널
	complete chan error
	// 지정된 시간이 초과했음을 알리는 채널
	timeout <-chan time.Time
	// 태스크를 위한 슬라이스
	tasks []func(int)
}
```

그리고 이 `Runner`를 생성하는 `New` 함수이다.

```go
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}
```

`time.Duration`을 매개 변수로 받는다. 실제로, `Runner`가 만들어질 때, 필드들은 다음과 같이 초기화된다.

**interrupt**

이 필드는 애플리케이션에서 인터럽트의 발생 유무를 확인한다. 한 번 발생하면 애플리케이션을 종료하므로, 버퍼의 크기가 1개인 채널을 할당한다.

**complete**

이 필드는 애플리케이션이 완료되었는지 아닌지 결과를 확인하는 채널이다.

**timeout**

이 필드는 `time.After` 함수를 통해서 `time.Time` 채널을 만든다. 실제 함수의 매개변수의 `d`만큼 시간이 지나면 이 채널에 메세지가 들어오게 되어 지정한 시간만큼 지났다는 것을 애플리케이션이 알 수 있게 해준다. 

tasks 필드는 `Nil` 슬라이스로 초기화한다. 이제 `Start` 메소드를 살펴보자.

```go
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}
```

먼저 `Runner`의 `interrupt`와 실제 OS에서 발생할 수 있는 인터럽트를 연결한다. 그 것이 `signal.Notify(r.interrupt, os.Interrupt)` 코드이다. 그리고 내부 메소드인 `run`을 고루틴으로 실행하여 그 결과를 `complete`에 넣어준다.

여기서 `select` 구문이 나온다. `select`는 채널에 `case`문에 걸어놓은 이벤트들이 발생할 때까지 대기한다. 만약 `timeout`에 메세지가 오면 시간이 지났기 때문에, 태스크 및 애플리케이션을 종료하기 위해 `ErrTimeout`을 반환한다. 이제 `complete`에 결과를 반환하는 `run`메소드를 살펴보자.


```go
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.goInterrupt() {
			return ErrInterrupt
		}

		task(id)
	}

	return nil
}
```

`run` 메소드는 `Runner`의 모든 태스크들을 순서대로 실행한다. 태스크 실행 전 `goInterrupt`를 호출한다. 그 후 태스크를 실행한다(`task(id)`). 

이제 `goInterrupt`를 보자. 

```go
func (r *Runner) goInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
```

위 코드는 `select` 구문을 활용하여, `interrupt` 채널에 수신된 인터럽트가 있는지 확인한다. 있으면, 그 인터럽트를 멈추고 true를 결과로 내보낸다. 없다면 false를 반환한다. 결국 `Runner`는 태스크를 순차적으로 실행하면서, 인터럽트가 발생하면, 인터럽트 예외를, 시간 초과를 발생하면, 시간 초과 예외를 발생시킨다. 참고적으로 아래 코드를 보자.

```go
var ErrTimeout = errors.New("시간을 초과했습니다.")
var ErrInterrupt = errors.New("OS 인터럽트 신호를 수신했습니다.")
```

이들은 패키지에서 발생할 수 있는 에러를 모아둔 것이다. `Go`에서는 에러를 정의할 때 이런 패턴을 사용한다. 잘 기억해두자. 이제 이 `Runner`를 활용한 main 코드를 살펴보자.

```go
package main

import (
	"log"
	"os"
	"time"

	"github.com/gurumee92/go-in-action/ch07/runner"
)

const timeout = 3 * time.Second

func main() {
	log.Println("작업 시작!")

	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrInterrupt:
			log.Println("인터럽트 에러 발생")
			os.Exit(2)
		case runner.ErrTimeout:
			log.Println("작업 시간 초과")
			os.Exit(1)
		}
	}

	log.Println("작업 종료!")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Process - JOB #%d Start.\n", id)
		time.Sleep(time.Duration(id) * time.Second)
		log.Printf("Process - JOB #%d Finish.\n", id)
	}
}
```

간단히 설명하면, `timeout`은 3초로 걸어둔다. 그리고 `Runner`를 생성하고, `Runner.run`에 전달할 id초만큼 실행 시 슬립되는 태스크를 3개 만들어둔다. 그리고 태스크들을 순차적으로 실행할 수 있도록 `Runner.Start`를 호출한다. 그러면, 태스크 결과에 따라 프로그램이 동작하게 된다. 그냥 애플리케이션을 실행할 경우 다음의 결과를 얻을 수 있다.

```
2020/09/22 23:15:43 작업 시작!
2020/09/22 23:15:43 Process - JOB #0 Start.
2020/09/22 23:15:43 Process - JOB #0 Finish.
2020/09/22 23:15:43 Process - JOB #1 Start.
2020/09/22 23:15:44 Process - JOB #1 Finish.
2020/09/22 23:15:44 Process - JOB #2 Start.
2020/09/22 23:15:46 작업 시간 초과
```

왜냐하면, 첫 번째 태스크는 0초 슬립, 두 번째 태스크는 1초 슬립 세 번째 태스크는 2초를 슬립한다. 그래서 실행 동작 시간까지 3초를 넘어가기 때문에, `timeout`에 걸려 강제 종료가 되는 것이다. 만약 애플리케이션 실행 중 강제 종료시키면 다음의 결과를 얻을 수 있다.

```
2020/09/22 23:17:37 작업 시작!
2020/09/22 23:17:37 Process - JOB #0 Start.
2020/09/22 23:17:37 Process - JOB #0 Finish.
2020/09/22 23:17:37 Process - JOB #1 Start.
^C2020/09/22 23:17:38 Process - JOB #1 Finish.
2020/09/22 23:17:38 인터럽트 에러 발생
exit status 2
```


## 재사용 가능한 리소스 풀 관리하기 Pool

이번에는, 버퍼가 있는 채널을 이용하여 공유가 가능한 리소스 풀을 생성하고 이를 관리하는 패턴을 알아보자. 이러한 패턴은 데이버테이스 연결이나, 메모리 버퍼 등 리소스의 정적인 집합을 관리할 때 매우 유용하다. 코드를 보자.

```go
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m sync.Mutex
	resources chan io.Closer
	factory func() (io.Closer, error)
	closed bool
}

var ErrPoolClosed = errors.New("Pool is closed.")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Pool's size too small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Resource acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Resource acquire:", "New Resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Resource Return: ", "Returned Resource Queue")
	default:
		log.Println("Resource Return: ", "Release Resource")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}
```

역시 하나 하나 뜯어보면서 보자. 먼저 `Pool` 구조체이다.

```go
type Pool struct {
	// 고루틴이 풀을 접근할 때, 안전하게 작업하기 위한 뮤텍스
	m sync.Mutex
	// io.Closer 채널, 버퍼가 있는 채널로 생성 리소스 공유가 목적
	resources chan io.Closer
	// 함수 타입.. 풀에 리소스 요청이 들어올 때 새로운 리소스 생성
	factory func() (io.Closer, error)
	// 풀이 닫혀있는지 확인하는 변수
	closed bool
}
```

`Pool`은 여러 개의 고루틴을 안전하게 공유하기 위한 리소스 집합을 표현한다. 리소스는 모두 `io.Closer`를 구현해야 한다. 이 구조체는 4개의 필드가 있는데, `m`은 뮤텍스 변수로 고루틴이 풀을 접근할 때 레이스 컨디션으로부터 안전하기 위해서 사용된다. `resources`는 버퍼가 있는 `io.Closer` 타입의 채널로써, 리소스 공유가 목적이다. `factory`는 풀에 리소스 요청이 들어올 때 새로운 리소스를 생성한다. `closed`는  `Pool`이 닫혀있는지 확인하기 위한 변수이다.

이제 `Pool`을 생성하는 `New` 함수를 보자.

```go
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Pool's size too small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}
```

먼저 매개 변수로, 입력은 x, 반환 타입이 `(io.Closer, error)` 인 함수 타입인 `fn`, 풀의 사이즈를 나타내는 `size`를 받는다. `size`가 0보다 작거나 같으면, 에러이다. 정상 입력이 되면, `factory`에 `fn`을, `resources`를 `io.Closer` 타입의 채널을 `size`만큼 할당해서 초기화한다.

이제 리소스를 획득하는 `Acquire` 메서드를 살펴보자.

```go
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Resource acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Resource acquire:", "New Resource")
		return p.factory()
	}
}
```

`Acquire` 메서드는 호출자가 풀에서 리소스를 획득할 수 있다. 채널에 리소스가 없으면, 새로운 리소스를 만들고 리소스가 있으면 그 리소스를 반환한다. 리소스가 만들어졌으면, 시스템에 다시 반환되어야 한다. 이 일을 하는 것이 `Release` 메서드이다.

```go
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Resource Return: ", "Returned Resource Queue")
	default:
		log.Println("Resource Return: ", "Release Resource")
		r.Close()
	}
}
```

이 메서드에서 사용하는 뮤텍스의 일은 크게 2가지이다.

1. `closed` 플래그를 읽을 때, `Close` 메서드가 이 플래그를 설정할 때 생기는 레이스 컨디션을 방지한다.
2. 닫힌 채널에 리소스를 돌려보내면 패닉이 발생한다. 이를 방지한다.

만약 풀이 닫혔을 경우는 해당 리소스를 아예 해제한다. 이제 풀을 아예 닫아버리는 `Close` 메서드를 보자.

```go
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}
```

여기서는 풀을 닫는다. `closed` 플래그를 바꾸고 `resources` 채널을 닫아버린다. 또한 `resources`에 있는 모든 리소스들을 닫아버린다. 여기서도 역시 뮤텍스 변수를 사용하여 경쟁 상태를 막는다. 이제 이를 활용하는 프로그램을 살펴보자.

```go
package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gurumee92/go-in-action/ch07/pool"
)

const (
	maxGoroutines   = 25
	pooledResources = 2
)

type dbConnection struct {
	ID int32
}

func (conn *dbConnection) Close() error {
	log.Println("Close: Database Connection ", conn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Database Connection ", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Program is done")
	p.Close()
}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("Query : QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
```

역시 하나하나 살펴보자.

```go
type dbConnection struct {
	ID int32
}

func (conn *dbConnection) Close() error {
	log.Println("Close: Database Connection ", conn.ID)
	return nil
}
```

먼저 데이터베이스 커넥션을 표현하기 위한 구조체와 해당 구조체의 `Close` 메서드이다. 


```go
var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Database Connection ", id)
	return &dbConnection{id}, nil
}
```

위의 코드는 고유한 ID를 만들기 위한 글로벌 변수와, 팩토리 함수이다. 이제 본격적인 메인 함수를 보자.

```go
func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("Program is done")
	p.Close()
}
```

먼저 고루틴 개수만큼 `sync.WaitGroup` 을 초기화한다. 그 후 `Pool`을 생헝한다. 이 때, 버퍼의 크기는 2, `dbConnection`을 만드는 팩토리 함수를 매개 변수로 전달한다. 그 후 원하는 개수만큼 고루틴을 생성한다.

고루틴에 싣는 `performQueries`를 보자.

```go
func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("Query : QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
```

풀에서 리소스를 획득하고, 이 함수가 끝나는 시점에 리소스를 해제한다. 여기서 일부러 슬립을 걸어준다. 쿼리가 실행되는 것을 흉내내는 것이다. 이제 애플리케이션을 실행해보자.

```
2020/09/23 21:57:26 Resource acquire: New Resource
2020/09/23 21:57:26 Create: New Database Connection  1
2020/09/23 21:57:26 Resource acquire: New Resource
...
2020/09/23 21:57:26 Query : QID[10] CID[24]
2020/09/23 21:57:26 Resource Return:  Returned Resource Queue
...
2020/09/23 21:57:27 Program is done
2020/09/23 21:57:27 Close: Database Connection  24
2020/09/23 21:57:27 Close: Database Connection  4
```

먼저 리소스 획득과 생성되는 것을 볼 수 있다. 25개가 만들어지면, Query가 실행 및 반환되는 것을 확인할 수 있다. 또한, 프로그램을 종료할 때 모든 리소스가 한꺼번에 반환되는 것을 볼 수 있다.


## 작업 처리 고루틴 풀 생성하기 Worker

이 패턴은 버퍼가 없는 채널을 이용하여, 원하는 개수만큼의 작업을 동시적으로 실행할 수 있는 고루틴 풀을 생성한다. 이 때는 크기가 있는 버퍼보다, 크기가 없는 버퍼를 사용한다. 그래야, 특정 작업 유실 없이 모든 작업을 동시적으로 실행할 수 있다. 코드를 보자.

```go
package work

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

func (p *Pool) Run(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
```

하나 하나 뜯어보자.

```go
type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}
```

`Worker`는 태스크를 실행하는 인터페이스이다. `Pool`은 `Worker` 타입의 크기가 없는 버퍼를 갖는 채널을 가지고 있다. 또한, 동시에 실행되는 개수를 조절하기 위하여 `sync.WaitGroup`을 사용한다. 이제 이 `Pool`을 생성하는 `New` 함수를 살펴보자.

```go
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}
```

일정한 개수의 고루틴을 생성하도록 작성된 작업 풀을 생성한다. 생성할 고루틴의 개수를 매개 변수로 전달되며, Pool 타입을 만드는 것을 볼 수 있다.

```go
for i := 0; i < maxGoroutines; i++ {
	go func() {
		for w := range p.work {
			w.Task()
		}
		p.wg.Done()
	}()
}
```

이 코드는 풀의 `work` 채널에서 `Worker` 인터페이스 값을 받는 한 계속 반복문을 실행한다. 루프 내에서 채널에서 받은 값에 대해 `Task`를 호출한다. 채널이 닫히면, 반복문의 실행이 종료되고, `WaitGroup`의 `Done` 메서드를 호출한 후, 고루틴 실행을 종료한다.

```go
func (p *Pool) Run(w Worker) {
	p.work <- w
}
```

이 메서드는 풀에 새로운 작업을 추가하는 코드이다. 풀의 `work` 채널의 매개 변수 `Worker`를 전달한다.

```go
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
```

이 메서드는 풀의 자원들을 반환한다. `work` 채널을 닫고 닫힐 때까지 `WaitGroup.Wait`를 호출하여 대기한다. 이제 메인 코드를 보자.

```go
package main

import (
	"log"
	"sync"
	"time"

	"github.com/gurumee92/go-in-action/ch07/work"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (p *namePrinter) Task() {
	log.Println(p.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(10)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	p.Shutdown()
}
```

또 코드를 하나하나 보자.

```go
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}
```

`names`는 출력할 이름들이 들어있다. 

```go
type namePrinter struct {
	name string
}

func (p *namePrinter) Task() {
	log.Println(p.name)
	time.Sleep(time.Second)
}
```

이 이름들을 `namePrinter`라는 구조체에 담는다. `namePrinter`는 `Worker` 인터페이스를 구현하기 때문에 `Task` 메서드를 가지고 있다. 그냥 이름을 출력하고 1초 정도 대기한다.

```go
func main() {
	p := work.New(10)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	p.Shutdown()
}
```

먼저 10은 동시에 실행될 고루틴의 개수를 지정한다. 그리고 이제 이름 개수 * 100을 `WaitGroup`을 초기화한다. 100번을 반복하면서, `names` 에 들어있는 이름을 필드로 갖는 `namePrinter`의 `Task`를 고루틴에 실어 보내며, 이는 `Pool`에 들어간다. 총 500번의 작업이 생성되는데 10번씩 동시에 출력된다. 그래서 50초면 작업이 끝나는 것을 확인할 수 있다.