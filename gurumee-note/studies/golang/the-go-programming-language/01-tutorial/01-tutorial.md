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

이 절에서는, 자료구조 `slice`, 반복문 `for`, 변수 선언 및 할당, 추가로 연산자에 대해 간단히 알아보자. 먼저 예시를 보자.

ch01/ex01_02_echo1.go
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
```

먼저 변수 선언을 살펴보자. 기본적으로 `Go`에서는 이렇게 변수 선언이 가능하다.

```
var 변수_이름 타입
```

위의 코드에서는 다음이 변수 선언 문이다.

```go
var s, sep string
```

`Go`는 변수에 값을 할당을 하지 않으면, 자동으로 기본 값을 지정한다. `string`타입의 경우, 자동 값은 `""`이다. 할당은 다음과 같이 할 수 있다.

```go
sep = " "
```

이를 한 번에 합치는 것도 가능하다.

```go
var sep = " "
```

이를 축약하는 것도 가능하다.

```go
sep := " "
i := 1
```

이것이 `Go`에서 변수 선언 및 할당하는 모든 구문이다. 이번엔 반복문 `for`와 슬라이스를 한꺼번에 보자.

```go
for i := 1; i < len(os.Args); i++ {
    s += sep + os.Args[i]
    sep = " "
}
```

먼저 반복문이다. `for` 키워드를 이용한다. 위는 가장 기본적인 형태의 반복문이다. 아래 형태를 띄고 있다.

```
for 변수 할당; 조건; 증감연산자 {
    반복 코드
}
```

반복문은 이 말고도 3가지 형태가 더 있다. 

```
for 조건 {
    반복 코드
}
```

이 형태로 변경해보자.

```go
i := 1

for i < len(os.Args) {
    s += sep + os.Args[i]
    sep = " "
    i += 1
}
```

이 말고 `for` 만 쓸 수도 있다.

```go
for {
    // ...
}
```

이 경우 `{}` 블록 안에서, 반복문을 탈출시키는 코드를 작성하지 않으면, 무한하게 반복한다. 또 한가지 특별한 형태가 있는데 이는 슬라이스를 설명하면서 같이 보자.

`os.Args`는 대표적인 슬라이스의 예제이다. 프로그램 실행 시 사용자가 입력한 매개 변수를 저장한다. 쉽게 생각하면 `가변 배열`이다. 파이썬의 슬라이스와 유사하다. 다만, `Go`는 포인터 기반이라, 얕은 복사/깊은 복사의 개념을 잘 생각해야 한다.

일반적으로 슬라이스가 저장하는 원소를 접근하기 위해서는 "인덱스"와 함께 접근할 수 있다.

```go
os.Args[1]
```

인덱스는 배열의 길이가 5라면 0~4까지를 나타낸다. 참고적으로, 배열 및 슬라이스의 길이를 알 수 있는 내장 함수 `len`을 제공한다.

```go
len(os.Args)
```

위의 코드는 `os.Args`의 길이를 반환한다. 또한, 슬라이스는 적절한 인덱스를 사용하면 자를 수가 있다. 

```
os.Args[1:]
```

이는 `os.Args` 인덱스 1에서 끝까지 자른다. 즉 원소 0에 대한 부분이 날라간다. 아까 `for`에 또 다른 형태가 있다고 했는데, 슬라이스, 맵 등의 자료구조는 `range` 키워드를 이용해서 순회할 수 있다.

```go
for idx, param := range os.Args[1:] {
    fmt.Println(idx, param)
}
```

`range`를 사용하면, 첫 번째로는 인덱스 요소를 반환한다. 두 번째로 실제 값을 반환하는데, 만약 `os.Args[1:]`이 다음과 같다고 해보자.

| 인덱스 | 값 |
| 0 | test0 |
| 1 | test1 |
| 2 | test2 |

그럼 결과는 다음과 같아진다.

```
0 test0
1 test1
2 test2
```

그럼 이제, 이를 이용해서, 위의 예제 코드를 바꿔보자. 변수 선언 및 할당, 반복문 for에 대해서 변경할 것이다.

ch01/ex01_03_echo2.go
```go
package main

import (
	"fmt"
	"os"
)

func main() {
    s, sep := "", ""
    
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
    }
    
	fmt.Println(s)
}
```


## 중복 줄 찾기

이 절에서는 제어문 `if`, 자료구조 `map`, 문자열 포매터, 그리고 `bufio`를 이용한 입/출력이나 `os`를 이용한 파일 입/출력에 대해서 살펴본다.

## 애니메이션 GIF

이 절에서는 `const` 선언, 구조체, 복합 리터럴 등에 대해서 살펴본다.

## URL 반입

이 절에서는 간단한 `net/http` 패키지에 대해서 다룬다.

## URL 동시 반입

이 절에서는 `Go`에서 지원하는 동시성 프로그래밍을 다루는 `go routine`과 `channel`에 대해서 다룬다.

## 웹 서버

이 절에서는 `net/http`를 이용한 웹 서버를 만드는 것에 초점을 둔다.