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


## 작업 처리 고루틴 풀 생성하기 Worker