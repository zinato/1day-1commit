# 표준 라이브러리

![대표사진](../logo.png)

> 책 "Go In Action"을 공부하면서 정리한 문서입니다.


## 표준 라이브러리를 공부하는 이유

`Go의 표준 라이브러리`눈 언어 자체를 향상시키고 확장시킬 수 있는 핵심 패키지들의 집합이다. 공부하는 이유는 `C++`의 `STL`, `Java`의 `Collection Framework`와 `Stream API` 같은 여러 표준 라이브러리들과 같다. 핵심 모토는 "바퀴를 재발명하지 말자."이다. 또한 책에서는 **하위 호환성**이 보장되므로 되도록 많이 사용하는 것을 권장하고 있다.

`Go`는 신생 언어답게 표준 라이브러리가 강력하며, 매우 방대한 생태계를 이루고 있다. 어지간히 큰 경우가 아니라면, 외부 라이브러리/프레임워크 없이 표준 라이브러리만으로도 해결이 가능하다. 이번 장에서는 `log`, `io`, `encoding` 등의 표준 라이브러리 패키지들을 살펴볼 것이다.

> 책에서는 문서화도 제공하고 있으나 개인적으로 기록할 필요는 없어 보인다.


## 로깅

`logging`은 버그를 찾거나, 프로그램이 어떻게 동작하는지 확인하는 아주 좋은 방법이다. `Go`에서는 `log` 패키지가 이를 지원한다. 먼저 다음의 `Go` 프로그램을 작성해보자.

```go
package main

import "log"

func init() {
	log.SetPrefix("Tracing: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("Message")
	log.Fatalln("Critical Error Message")
	log.Panicln("Panic Message")
}
```

위의 프로그램은 먼저, `main`을 호출하기 이전에 `init` 함수를 호출하여 로그를 설정한다. 로그 앞에 "Tracing: " 이란 문자열을 붙이고, 날짜, ms, 파일 명을 표시한다. 위의 프로그램을 실행하면 다음 결과를 얻을 수 있다.

```bash
$ go run ch08/example_log_01.go
Tracing: 2020/10/01 21:40:34.842233 /Users/gurumee/Studies/go-in-action/ch08/example_log_01.go:11: Message
Tracing: 2020/10/01 21:40:34.842338 /Users/gurumee/Studies/go-in-action/ch08/example_log_01.go:12: Critical Error Message
exit status 1
```

출력 구조가 로그를 설정한 값을 따르고 있는 것을 볼 수 있다. 세 번째 `Panicln`이 호출이 되지 않는 것은 `Fatalln`이 호출되면, 프로그램이 강제로 종료되기 때문이다. 이는 `Fatalln`과 `Panicln`을 순서를 바꾸어도 마찬가지다. 

로그의 플래그 값들은 `Go` 공식 문서에서 확인할 수 있다.

```go
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
```

더 자세한 내용을 알고 싶다면, [이 곳](https://pkg.go.dev/log)을 클릭하라. 또한 다음과 같이 사용자 정의 로거도 만들 수 있다.

```go
package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("에러 로그 파일을 열 수 없습니다.", err)
	}

	Trace = log.New(ioutil.Discard, "Tracing: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("일반적인 로그 메시지")
	Info.Println("특별한 정보를 위한 로그 메시지")
	Warning.Println("경고성 로그 메시지")
	Error.Println("에러 로그 메시지")
}
```

`log.New`로 새로운 로거를 만들 수 있다. 굉장히 흥미로운데, 첫 번째 인수는 출력 스트림을 지정한다. `ioutil.Discard`는 아무 곳도 지정하지 않는다. 따라서 출력되지 않는다. `os.Stdout`은 표준 출력 스트림을 지정한다. 그럼 터미널에 메세지가 출력된다. 또한, `io.MultiWriter`는 여러 출력 스트림을 지정할 수 있다. `Error` 로그는 `os.Stderr` 표준 에러 스트림과, `error.txt` 파일에 로그를 출력시킨다.

실제 프로그램을 돌려보면 다음과 같다.

```bash
go run ch08/example_log_02.go
Info: 2020/10/01 21:50:14 example_log_02.go:32: 특별한 정보를 위한 로그 메시지
Warning: 2020/10/01 21:50:14 example_log_02.go:33: 경고성 로그 메시지
Error: 2020/10/01 21:50:14 example_log_02.go:34: 에러 로그 메시지
```

그리고, `error.txt`에서는 `Error` 로그 출력을 확인할 수 있다.

```txt
Error: 2020/10/01 21:54:03 example_log_02.go:34: 에러 로그 메시지
```

`Go`의 `log` 패키지는 이미 많은 개발자들이 사용하고, 성숙한만큼 성숙했다. 즉 신뢰 가능한 패키지라는 말이다. 애플리케이션을 만들 때 로그가 필요하다면, 이를 기억하자.


## 인코딩과 디코딩

## 입력과 출력