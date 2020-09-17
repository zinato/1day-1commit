# 동시성

![대표사진](../logo.png)

> 책 "Go In Action"을 공부하면서 정리한 문서입니다.



## 프로세스와 스레드, 동시성과 병렬성

`Go`에서의 동시성 처리 동기화는 `CSP`라는 패러다임에서 비롯되었다. `CSP`는 간단하게 말해, 동시 접근에 대해 잠금으로 처리하는 것이 아닌, 메세지를 전달하는 방식이다. 이 메세지를 전달하는 통로가 바로 "채널"이다.

![프로세스](./01.png)

위의 그림은 **프로세스**의 모습이다. 메모리 주소 공간, 파일 및 장치, 스레드에 대한 핸들을 비롯해 다양한 것이 포함된다. **스레드는 코드로 작성한 함수를 실행하기 위해 운영체제가 예약해둔 실행 경로**이다. 프로세스는 최소 하나의 스레드를 지닌다. (스레드는 프로세스의 포함됨) 
 

![런타임 스케줄러 일반적인 관리](./02.png)

위의 그림은, `Go 런타임 스케줄러`는 하나의 운영체제 스레드에 바인딩된 논리적 프로세서에서 고루틴이 실행될 수 있게 예약함을 보여준다. 보통 다음의 순서를 따른다.

1) 고루틴이 생성된다.
2) 고루틴이 실행할 준비가 되면 스케줄러의 `Global Run Queue`에 위치한다.
3) 논리 프로세서가 고루틴을 큐에서 빼와서 실행한다.

![런타임 스케줄러 시스템 콜 발생 시 관리](./03.png)

위의 그림은 파일을 여는 등의 자신의 실행을 중단해야 하는 시스템 콜을 수행하는 경우를 보여준다. 다음의 순서를 따른다.

1) 논리 프로세서와 고루틴이 분리된다.
2) 스레드는 시스템 콜이 리턴될 때까지 대기한다. 
3) 논리 프로세서는 할당된 스레드가 없기 때문에, 스케줄러에서 새로운 스레드를 생성한다.
4) 새로운 스레드를 논리 프로세스에 연결한 후 `Local Run Queue`에서 다른 고루틴을 선택하여 실행한다.
5) 시스템 콜이 리턴되면 실행 중인 고루틴은 다시 `Local Run Queue`로 이동한다.

5번이 실행될 때, 다시 사용될 것을 대비하여 고루틴이 실행 중이던 스레드도 함꼐 보관된다.

![동시성과 병렬성](./04.png)

위의 그림은 동시성과 병렬성의 차이를 보여준다. 이 둘은 다른 개념이다. 병렬성은 여러 논리 프로세스서에서 동시에 실행된다. 동시성은 한 번에 여러 작업을 수행하는 것을 말한다. 

보통 운영체제와 하드웨어에 가해지는 부담이 적기 때문에 동시성 처리가 병렬성 처리보다 성능이 우수한 경우가 많다.


## 고루틴

`고루틴`은 일종의 경량 스레드이다. 쓰는 방법은 쉽다. 함수 호출 앞에 `go` 라는 키워드만 붙이면 된다.

다음 예제를 보자.

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 스케줄러가 사용할 하나의 논리 프로세스 할당
	runtime.GOMAXPROCS(1)

	// wg는 프로그램 종료를 대기하기 위해 사용
	// 고루틴 개수만큼 더해준다.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("고루틴 시작!")

	go func() {
		// main 함수에 종료를 알리기 위해 Done 함수 호출
		defer wg.Done()

		// 소문자를 3번 출력한다.
		for count := 0; count < 3; count++ {
			time.Sleep(3 * time.Second)
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	go func() {
		// main 함수에 종료를 알리기 위해 Done 함수 호출
		defer wg.Done()

		// 대문자를 3번 출력한다.
		for count := 0; count < 3; count++ {
			time.Sleep(3 * time.Second)
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	fmt.Println("고루틴 대기 중")
	wg.Wait()

	fmt.Println("고루틴 끝~!")
}
```

한 코드 한 코드 뜯어서 보자.

```go
runtime.GOMAXPROCS(1)
```

위 코드는 주석에도 적혀져 있지만, 논리 프로세스의 개수를 조절하는 것이다. 따라서 이 프로그램이 실행되면 단 하나의 논리 프로세스가 할당된다.

```go
var wg sync.WaitGroup
wg.Add(2)

fmt.Println("고루틴 시작!")
// ...


fmt.Println("고루틴 대기 중")
wg.Wait()

fmt.Println("고루틴 끝~!")
```

`WaitGroup`은 "카운팅 세마포어"라는 방식으로 고루틴이 종료될 때까지 메인 스레드가 유지되게끔 한다. 고루틴의 개수만큼 `WaitGroup`에 추가해주어야 하며, `Wait`를 호출해야 설정한 개수만큼 고루틴이 끝날 때까지 대기한다. 

```go
// ...
go func() {
		// main 함수에 종료를 알리기 위해 Done 함수 호출
    defer wg.Done()

    // 소문자를 3번 출력한다.
    for count := 0; count < 3; count++ {
        time.Sleep(3 * time.Second)
        for char := 'a'; char < 'a'+26; char++ {
            fmt.Printf("%c ", char)
        }
        fmt.Println()
    }
}()
// ...
```

`defer`를 사용해서 고루틴이 종료 시에 `WaitGroup.Done`을 호출하여 고루틴이 끝났음을 알린다. 이는 익명 함수로써 소문자를 3번 출력하고 끝이 난다. 아래 익명함수는 대문자를 출력하고 끝이난다.

프로그램을 실행해보면, 3초마다 대문자/소문자 출력이 무작위로 출력되는 것을 볼 수 있다.

```
A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 
a b c d e f g h i j k l m n o p q r s t u v w x y z 
a b c d e f g h i j k l m n o p q r s t u v w x y z 
A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 
A B C D E F G H I J K L M N O P Q R S T U V W X Y Z 
a b c d e f g h i j k l m n o p q r s t u v w x y z 
```


## 경쟁 상태와 락 기법

이런 동시성 처리를 할 때 주의할 점이 있다. 스레드끼리는 자원을 공유하기 때문에, 일종의 `경쟁 상태(Race Condition)`라고 불리우는 골치 아픈 문제가 발생한다. 아래 예제를 보자.

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func intCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter
		runtime.Gosched()
		value++
		counter = value
	}
}

func main() {
	wg.Add(2)

	go intCounter(1)
	go intCounter(2)

	wg.Wait()
	fmt.Println("Result: ", counter)
}
```

쉽게 설명하면, 고루틴 2개가 번갈아가면서 실행되면서 `counter`의 값을 각각 2번씩 반복하면서 1씩 증가시키는 프로그램이다. 그런데 코드를 실행해보면 결과는 2가 나올 때가 있다. (4, 2 중 하나가 나온다.) 

![경쟁 상태](./05.png)

위의 그림은 앞선 예제에서 `counter`가 고루틴 2개에 의해서 덮어 씌어지는 것을 보여준다. 

`Go`는 `go run` 혹은 `go build` 명령 시에 이런 경쟁 상태를 트레이싱할 수 있는 옵션을 제공한다. `-race`를 붙이면 된다.

```bash
# go run -race 파일
$ go run -race example.go
==================
WARNING: DATA RACE
Read at 0x00000127f540 by goroutine 8:
  main.intCounter()
      /Users/gurumee/Studies/go-in-action/ch06/example_06_02_race_condition.go:18 +0x79

Previous write at 0x00000127f540 by goroutine 7:
  main.intCounter()
      /Users/gurumee/Studies/go-in-action/ch06/example_06_02_race_condition.go:21 +0x9a

Goroutine 8 (running) created at:
  main.main()
      /Users/gurumee/Studies/go-in-action/ch06/example_06_02_race_condition.go:29 +0x89

Goroutine 7 (finished) created at:
  main.main()
      /Users/gurumee/Studies/go-in-action/ch06/example_06_02_race_condition.go:28 +0x68
==================
Result:  4
Found 1 data race(s)
exit status 66
```

실제 실행 시 레이스 컨디션을 트레이싱 하게 해보았다. 예상대로 `intCounter` 함수에서 레이싱 컨디션이 발생할 수 있다고 경로를 띄우고 있다.

이런 해결책으로 "공유 자원 기법(Lock 기법)"이라는 방법이 있다. 먼저 `atomic`을 사용하는 것이다. 다음 코드를 보자.

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func intCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}

func main() {
	wg.Add(2)

	go intCounter(1)
	go intCounter(2)

	wg.Wait()
	fmt.Println("Result: ", counter)
}
```

`atomic`은 정수 및 포인터에 대한 접근을 동기화할 수 있는 저수준의 잠금 매커니즘을 제공한다.

위 코드처럼 경쟁 상태에 놓여 있는 변수에 일종의 락을 걸어서 연산이 끝난 후 락을 해제한다. 그래서 경쟁 상태를 발생하지 않게 한다. 실제 위의 파일을 `-race` 옵션으로 실행해주면 해당 경고들이 사라짐을 볼 수 있다.

```bash
go run -race example.go
Result:  4
```

다른 잠금 기법은 `뮤텍스`를 활용하는 것이다. 이는 `상호 배타(mutual exclusion` 개념을 추상화한 것이다. 어떤 임계 지역을 생성하여 이 지역에는 하나의 고루틴만 통과할 수 있게 만들어두는 것이다. 코드를 보자.

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func intCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}
}

func main() {
	wg.Add(2)

	go intCounter(1)
	go intCounter(2)

	wg.Wait()
	fmt.Println("Result: ", counter)
}
```

`intCounter`를 자세히 보자.


```go
func intCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
        // 임계 부분 시작 점
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
        }
        // 임계 부분 끝 점
		mutex.Unlock()
	}
}
```

`mutex.Lock` 이후 부터는 임계 지역에 들어간다. `mutex.Unlock` 호출 이후 임계 지역이 풀린다. 물론 이런 방법들이 훌륭하긴 하지만 여전히 어렵다. 

잠금 기법의 경우, 잠금을 적절히 해제하지 않으면 "데드락"이라는 심각한 문제를 초래하게 된다. 그래서 조심 또 조심해서 코딩을 해야 한다.


## 채널

버퍼가 없는 채널

버퍼가 있는 채널