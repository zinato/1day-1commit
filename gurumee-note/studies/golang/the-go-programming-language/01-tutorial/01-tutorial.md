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

연습 문제 1.1
```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := strings.Join(os.Args, " ")
	fmt.Println(s)
}
```

연습 문제 1.2
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[1:] {
		fmt.Printf("idx: %v, value: %v\n", idx, arg)
	}
}
```

연습 문제 1.3
```go
package test

import (
	"fmt"
	"os"
	"testing"
)

func echo1() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	for idx, arg := range os.Args[1:] {
		fmt.Printf("idx: %v, value: %v\n", idx, arg)
	}
}

func BenchmarkEcho1(b *testing.B) {
	b.ResetTimer()
	os.Args = []string{"echo1", "hello", "world"}

	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	b.ResetTimer()
	os.Args = []string{"echo2", "hello", "world"}

	for i := 0; i < b.N; i++ {
		echo2()
	}
}
```


## 중복 줄 찾기

이 절에서는 제어문 `if`, 자료구조 `map`에 대해서 살펴본다. 바로 예를 살펴보자.

ch01/ex01_05_dup1.go
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		s := input.Text()

		if s == "" {
			break
		}

		counts[s]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
```

`map`은 키-값 쌍으로 데이터를 저장하는 자료구조이다. `Go`에서 기본적으로 제공하는 복합 타입 중 하나이다. `map`은 `make` 내장 함수로 만들 수 있다. 다음과 같이 말이다.

```go
counts := make(map[string]int)
```

이렇게 하면 문자열 타입을 키, 정수형 타입을 값 쌍으로 저장하게 된다. 데이터를 저장하려면 이렇게 할 수 있다.

```go
counts["A"] = 5
```

그러면, "A"라는 키를 가지고 값을 가져오려면 어떻게 해야 할까? 똑같이 `[]` 연산자를 쓰면 된다. 이 때 주의할 점은 이 연산자는 2개의 반환값을 내놓은다.

```go
count, exists := counts["A"]
```

이 경우, 키 "A"에 쌍으로 저장되어있는 값이 있다면, `값, true` 값이 없다면, `nil, false`가 반환하게 된다. 즉 위의 코드에서는 `count=5, exists=true`가 반환되어 각각 저장된다. 키를 `map`에서 지우고 싶다면, 다음 `delete` 내장 함수를 쓰면 된다.

```go
delete(counts, "A")
```

또한, `map` 역시 `slice`처럼, `for`문으로 순회할 수가 있다. 다음과 같이 말이다.

```go
for k, v := range counts {
	// ...
}
```

이러면, 한 번 순회할 때마다, `k`에는 키가, `v`에는 값이 할당되어 사용할 수 있다. 이제 `if`문을 보자.

```go
for line, n := range counts {
	if n > 1 {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
```

여기서는 중복된 단어의 횟수가 여러 개의 경우, 즉 1개보다 많은 경우 단어를 출력하게 만드는 것이다. `if`는 제어문이다. 조건에 따라, 분기 처리를 위해서 쓴다. 다음과 같은 구조이다.

```
if 조건문 {
    // 조건문이 만족하면, 실행
}
```

그리고, 조건문에 만족하지 않을 때 나누고 싶다면, `else`를 쓰면 된다.

```
if 조건문 {
    // 조건문이 만족하면, 실행
} else {
	// 조건문이 만족하지 않으면, 실행
}
```

또한, 여러 조건을 비교해서 실행하려면, `if - else if - else` 구조로 만들 수 있다.

```
if 조건문 1 {
    // 조건문 1이 만족하면, 실행
} else if 조건문 2 {
	// 조건문 1이 만족하지 않고, 조건문 2를 만족하면 실행
} else {
	// 조건문 1, 2를 만족하지 않으면 실행
}
```

책에서는 `bufio`를 이용해서, 표준 입출력 `os.Stdin`을 이용하는 부분이 나오는데, 잠깐 살펴보자.

```go
input := bufio.NewScanner(os.Stdin)
```

`input`은 `Go`의 표준 입력을 읽어서, 줄 혹은 단어 단위로 나누는 `Scanner` 타입이다. `bufio` 패키지에 정의되어 있다. 표준 입력으로 들어오는 문자열 스트림을 처리할 수 있다. 

```go
for input.Scan() {
	s := input.Text()

	if s == "" {
		break
	}

	counts[s]++
}
```

위 코드는 반복적으로 `input.Scan` 메소드를 호출하여, 터미널에 입력한 문자열("Enter 입력 전까지")을 읽는다. 그래서, 해당 문자열을 반환하는 것은 `input.Text` 메소드이다. 만약 아무 문자열을 입력 안했을 때, `break` 문을 이용해서 반복문을 멈춘다.


연습 문제 1.4
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang-collections/collections/set"
)

func main() {
	counts := make(map[string]*set.Set)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			log.Fatalln(err)
		}

		for _, line := range strings.Split(string(data), "\n") {
			if counts[line] == nil {
				counts[line] = set.New()
			}

			counts[line].Insert(filename)
		}
	}

	for line, sets := range counts {
		if sets.Len() > 1 {
			fmt.Printf("%v: ", line)
			sets.Do(func(i interface{}) {
				fmt.Printf("%v ", i)
			})
			fmt.Println()
		}
	}
}
```


## 애니메이션 GIF

이 절에서는 `const` 선언, 구조체, 복합 리터럴 등에 대해서 살펴본다. 다음은 예제 코드이다.

ch01/ex01_08_lissajous.go
```go
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
)

var palette = []color.Color{color.White, color.Black}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
```

꽤 긴데, 다루는 내용은 짧다. 먼저 `const` 키워드이다. `Go`에서 변하지 않는 수, 즉 상수를 선언 및 할당을 위해서는 이 키워드를 사용해야 한다.

```go
const (
	cycles  = 5
	res     = 0.001
	size    = 100
	nframes = 64
	delay   = 8
)
```

위의 코드에서 `cycles`, `res`, `size`, `nframes`, `delay` 모두 상수이다. 이들은 한 번 초기화가 이루어지면 프로그램이 실행해서 종료할 때까지 바뀌지 않는다. 만약 `const` 선언 후 초기화가 일어나지 않는다면 어떻게 될까?

```go
const test
```

그러면 컴파일 에러가 뜬다. 상수 선언 후 반드시 초기화까지 이뤄져야 한다. 이제 구조체를 살펴보자. 구조체는, 여러 타입을 하나의 타입으로 묶어주는 역할을 한다. 뭐 예를 들어 사람이라면, 이름과 나이를 가져야 한다고 가정해보자. 그럼 `Go`로 어떻게 표현할까? 

```go
type Person struct {
	Name string
	Age  int
}
```

위 코드는 구조체를 선언하는 코드이다. `Person`은 구조체 이름이다. 또한 이름을 표현하는 `Name`, 나이를 표현하는 `Age`는 구조체의 필드라고 부른다. 위의 예제에서는 `gif.GIF`가 바로 구조체이다.

```go
anim := gif.GIF{LoopCount: nframes}
```

이는 `gif.GIF` 구조체의 객체를 하나 생성하는 것이다. 내부 필드로 `LoopCount`를 가지고 있는데, `nframes`로 값을 할당하는 것이다. 실제 구조체는 다음과 같다.

```go
type GIF struct {
	Image []*image.Paletted 
	Delay []int             
	LoopCount int
	Disposal []byte
	Config image.Config
	BackgroundIndex byte
}
```

`LoopCount`를 제외한 나머지 필드들은 어떻게 될까? 각 타입의 기본 값으로 초기화 된다.


연습 문제 1.5
```go
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0, 230, 64, 1}}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
```

연습 문제 1.6
```go
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{255, 0, 0, 1},
	color.RGBA{0, 255, 0, 1},
	color.RGBA{0, 0, 255, 1},
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	index := 0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(index))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

		index = (index + 1) % len(palette)
	}

	gif.EncodeAll(out, &anim)
}
```


## URL 반입

이 절에서는 간단한 `net/http` 패키지에 대해서 다룬다. 또한, `defer`에 대해서 다룬다. 다음 예제를 살펴보자.

ch01/ex01_09_fetch.go
```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%s", b)
	}
}
```

먼저, `http` 메소드에는 여러 개 있다. `GET`, `POST`, `PUT`, `DELETE` 등이 있는데 이에 대한 `http` 콜을 `net/http` 패키지에 구현되어 있다. 여기서는 url에 대해서 `GET` 방식으로 Call을 하게 된다.

```go
defer resp.Body.Close()
```

위 코드는 `defer`를 사용한 것이다. `defer`는 함수가 종료될 때 리소스를 해제시키고 싶을 때 쓰는 키워드이다. 이렇게 하면, `resp.Body.Close()`는 함수 종료 직전에 무조건 호출하게 된다. 뭐 `resp.Body`는 출력 스트림인데, `ioutil.ReadAll` 함수로, `[]bytes` 타입으로 바꿀 수 있다. 또 이는 `string` 함수를 통해 문자열로 만들어줄 수 있다.

연습 문제 1.7
```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		defer resp.Body.Close()
		_, err = io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
```

연습 문제 1.8
```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		defer resp.Body.Close()
		_, err = io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
```

연습 문제 1.9
```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		defer resp.Body.Close()
		_, err = io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("Status Code: %v\n", resp.Status)
	}
}
```

## URL 동시 반입

이 절에서는 `Go`에서 지원하는 동시성 프로그래밍을 다루는 `go routine`과 `channel`에 대해서 다룬다. **이 절이 하이라이트이다.** 먼저 예제를 보자.

ch01/ex01_10_fetchall.go
```go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	defer resp.Body.Close()
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2ffs %7d %s", secs, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
```

`go routine`은 `Go`에서 만드는 경량 스레드이다. 단지 `go`라는 키워드를 함수 호출 시에, 같이 붙이기만 해도, 고루틴이 만들어져서 동시적으로 호출된다.

```go
for _, url := range os.Args[1:] {
	go fetch(url, ch)
}
```

이거는 프로그램 인수로 전달된 URL 개수만큼 `fetch` 함수에 대한 고루틴을 만들어서 동시 호출하게 된다. URL이 3개라면, 3개가 동시에 호출된다. `channel`은 `go routine` 사이의 통신을 가능하게 해준다. 타입은 `chan`이고 `make` 함수로 만들 수 있다.

```go
ch := make(chan string)
```

이것은 문자열 채널을 만드는 것이다. 고루틴의 결과를 문자열로 채널에 전달할 수 있다. `fetch` 함수를 보라.

```go
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	defer resp.Body.Close()
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2ffs %7d %s", secs, nbytes, url)
}
```

먼저 함수 파라미터를 보자.

```go
func fetch(url string, ch chan<- string) {
	// ...
}
```

`ch chan<- string` 이것은 고루틴으로부터 결과만 입력 받는 채널을 명시한다. 채널에 문자열 결과를 전달하기 위해서는 다음처럼 하면 된다.

```go
ch <- fmt.Sprintf("%.2ffs %7d %s", secs, nbytes, url)
```

그리고 이제 채널에서 결과를 꺼내려면 어떻게 해야할까? 다시 main 함수를 보자.

```go
func main() {
	// ...

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		result := <-ch
		fmt.Println(result)
	}

	// ...
}
```

먼저 채널에서 고루틴으로부터의 결과를 얻기 위해서는 다음처럼 쓰면 된다.

```go
result := <-ch
```

이 때 채널에서 결과를 꺼내는 것은 동기적으로 호출된다. 위의 코드는 `fetch`를 URL 개수만큼, 고루틴을 만들어 동시 호출한다. 그 후, 만들어진 고루틴 개수만큼 고루틴이 종료되서 결과를 전달받을 때까지 대기하고 출력하는 코드이다.

연습 문제 1.10
```go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	defer resp.Body.Close()
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main() {
	f, err := os.Create("practice_10.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Fprintln(f, <-ch)
	}

	fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds())
}
```

연습 문제 1.11
코드 아님.


## 웹 서버

이 절에서는 `net/http`를 이용한 웹 서버를 만드는 것에 초점을 둔다. `Go`는 웹 서버를 쉽게 만들 수 있다. 예제를 보자.

ch01/ex01_11_server1.go
```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
```

대단하지 않은가? 고작 몇 줄의 코드로 웹 서버를 만들 수 있다. 이는 개인적으로 너무 범위를 벗어나는 것 같다. 책 예제에서는 `sync.Mutex`를 활용하여, 웹 서버에서 발생하는 레이스 컨디션을 해결하는 방법을 설명하고 있는데 이도 생략한다.

연습 문제 1.12
```go
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{255, 0, 0, 1},
	color.RGBA{0, 255, 0, 1},
	color.RGBA{0, 0, 255, 1},
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	index := 0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(index))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

		index = (index + 1) % len(palette)
	}

	gif.EncodeAll(out, &anim)
}
```