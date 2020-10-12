# 튜토리얼

![로고](../logo.png)

> 책 "The Go Programming Language"를 읽고 정리한 문서입니다. 소스 코드는 다음에서 찾아볼 수 있습니다.
> 
> https://github.com/gurumee92/golang-studies/tree/the-go-programming-language/the-go-programming-language/ch01


## Hello, World

이 절에서는 `Go` 소스 코드의 기본 구조에 대해서 살펴본다. `package`, `import`, `func` `main`이 바로 그들이다.

ch01/ex01_01_helloworld.go
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

`Go`는 컴파일 언어이다. 즉, `Go`가 설치된 머신에 맞게, 바이너리 파일로 변환할 수 있다. `Go` 코드를 바이너리로 파일로 빌드하기 위해서는 `go build` 명령어를 이용한다. 다음처럼 말이다.

```bash
$ go build ch01/ex01_01_helloworld.go
```

그럼, 명령어를 실행한 위치에 `ex01_01_helloworld` 실행 파일이 만들어진다. 이제 이를 실행시키고 지운다.

```bash
# 실행
$ ./ex01_01_helloworld
Hello World

$ rm ex01_01_helloworld
```

코드 빌드, 실행, 삭제를 한 번에 해주는 `go run` 명령어도 있다. 터미널에 다음을 입력한다.

```bash
$ go run ch01/ex01_01_helloworld.go
Hello World
```

이제 코드를 하나 하나 뜯어보자.

```go
package main
```

이 구문은 `package` 선언 문이다. `package main` 선언은 조금 특별하다. 이렇게 선언을 한 소스파일만 빌드해서, 바이너리 파일로 만들 수 있다. 즉, `package main` 선언 한 곳은 `main` 함수가 존재해야 하며, 해당 소스 코드는 실행할 수 있다.

```go
import "fmt"
```

이 구문은 `fmt` 패키지를 소스 코드 내에 임포트한 것이다. `main` 함수에서 `fmt.Println()`을 사용하기 위해 반드시 선언해야 한다.

패키지는 패키지의 역할을 정의하는 하나의 디렉토리와, `.go` 확장자를 갖는 여러 소스 파일을 가지고 있다. `Go`에서 표준 입/출력을 담당하는 위의 `fmt` 패키지가 그 예이다.

`fmt` 패키지는 다음의 형태를 띈다.

```
|- fmt
    |- doc.go
    |- errors.go
    |- errors_test.go
    |- example_test.go
    |- export_test.go
    |- fmt_test.go
    |- foramt.go
    |- gostringer_example_test.go
    |- print.go
    |- scan.go
    |- scan_test.go
    |- stringer_example_test.go
    |- stringer_test.go
```

여기서 `print.go`의 패키지 선언을 살펴본다.

src/fmt/print.go
```go
package fmt

import (
	"internal/fmtsort"
	"io"
	"os"
	"reflect"
	"sync"
	"unicode/utf8"
)


// ...

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}

// ...
```

결국 아래에 쓰이는 `fmt.Println`은 위의 함수를 호출한 것이다. 여기서 패키지 외부로 노출시키려면 위처럼 대문자로 시작해야 한다. 패키지 안에서만 사용한다면, 소문자이다. 그리고 호출 시, `패키지_명.함수_명`의 형식으로 호출할 수 있다.

```go
func main() {
    fmt.Println("Hello World")
}
```

`main` 함수는 프로그램의 진입점이다. 오직 `main` 함수를 가지고 있는 소스파일만이, 빌드되어 실행 파일이 될 수 있다. 위에서 얘기한 것처럼 이 소스 코드는 "Hello World"만 출력하고 프로그램을 종료한다. 여기서 알 수 있는 것은 함수 선언 및 정의이다.

```
func 함수_이름(매개_변수_이름 타입, ...) (타입) {
    ...
}
```

기본적으로 `Go`는 `func` 키워드를 이용하여, 함수 선언 및 정의를 한다. 함수 이름과 함께, 매개변수 목록들을 이름과, 타입으로 선언해주어야 한다. 또한 뒤에는, 반환 타입을 선언해주어야 한다. 예를 들어서 이렇게 말이다.

```go
func Add(a, b int) int {
    return a + b
}
```

참고적으로, `go fmt` 명령어를 이용하면, 소스 코드를 정렬할 수 있다.


## 커맨드 라인 인수

## 중복 줄 찾기

## 애니메이션 GIF

## URL 반입

## URL 동시 반입

## 웹 서버