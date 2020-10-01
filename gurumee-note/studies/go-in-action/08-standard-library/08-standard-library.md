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

현대의 웹 개발 시, `API`와 `APP`을 분리한다. 뭐 더 나눠볼 순 있겠지만, 크게 이 2가지로 나눠보자. 이 때 동신은 무엇으로 할까? 보통 쉽게 떠올리는 것이 `JSON`이다. 예전엔, (현대에도 많이 사용하지만) `XML`도 사용했었다. `encoding` 패키지는 `JSON`, `XML`등의 데이터를 파싱하는데, 그 목적이 있다.

다음 프로그램을 만들어보자.

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Points        int    `json:"points"`
	User          string `json:"user"`
	Time          int    `json:"time"`
	TimeAgo       string `json:"time_ago"`
	CommentsCount int    `json:"comments_count"`
	Type          string `json:"type"`
	Url           string `json:"url"`
	Domain        string `json:"domain"`
}

func main() {
	uri := "https://api.hnpwa.com/v0/news/1.json"
	resp, err := http.Get(uri)

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	defer resp.Body.Close()

	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	pretty, err := json.MarshalIndent(posts, "", "    ")

	if err != nil {
		log.Fatalln("Error: ", err)
	}

	fmt.Println(string(pretty))
}
```

"https://api.hnpwa.com/v0/news/1.json" API에서 제공하는 데이터의 형태는 다음과 같다.

```
[
    {
        id: 24649992,
        title: "Programming Language notation is a Barrier to Entry",
        points: 22,
        user: "seg_lol",
        time: 1601554352,
        time_ago: "an hour ago",
        comments_count: 23,
        type: "link",
        url: "https://blog.sigplan.org/2020/09/29/pl-notation-is-a-barrier-to-entry/",
        domain: "blog.sigplan.org"
    },
    // ....
]
```

이 JSON 객체를 표현한 것이 바로 `Post`이다.

```go
type Post struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Points        int    `json:"points"`
	User          string `json:"user"`
	Time          int    `json:"time"`
	TimeAgo       string `json:"time_ago"`
	CommentsCount int    `json:"comments_count"`
	Type          string `json:"type"`
	Url           string `json:"url"`
	Domain        string `json:"domain"`
}
```

`json:"JSON 필드 이름"`을 지정하면 된다. 이제 디코딩, 인코딩 부분인데, 먼저 디코딩 부분이다.

```go
var posts []Post
err = json.NewDecoder(resp.Body).Decode(&posts)

if err != nil {
    log.Fatalln("Error: ", err)
}
```

먼저 디코딩은 `JSON` 객체를 실제 `Post` 같은 우리가 사용할 구조체를 옮기는 작업이라고 생각하면 편하다. 디코딩할 객체 `resp.Body`를 디코더에 넣는다. 그리고 `Decoder.Decode`의 해당 디코딩해서 값을 저장할 객체 `posts`의 주소를 넣어준다. 응답 `JSON`이 `Post` 리스트이기 때문에, `Post` 슬라이스를 넣어주면 된다.

그 다음 인코딩이다. 디코딩의 반대이다. 프로그램 객체를 `JSON` 객체로 바꿔준다. 다음은 `JSON` 객체를 이쁘게 출력하는 방법이다.

```go
pretty, err := json.MarshalIndent(posts, "", "    ")

if err != nil {
    log.Fatalln("Error: ", err)
}

fmt.Println(string(pretty))
```

실제 프로그램을 실행하면 다음의 결과를 얻을 수 있다.

```bash
[
    // ...
    ,
    {
        "id": 24638438,
        "title": "Gitter is joining Matrix",
        "points": 642,
        "user": "BubuIIC",
        "time": 1601472107,
        "time_ago": "a day ago",
        "comments_count": 108,
        "type": "link",
        "url": "https://matrix.org/blog/2020/09/30/welcoming-gitter-to-matrix",
        "domain": "matrix.org"
    }
]
```

알아둘 점은 `JSON` 안에 `JSON`이 겹친다면, 구초제도 겹치는 구조로 선언하면 된다. 또한, 유연성을 갖추기 위해서, `map[string]interface{}`로 선언해야 할 때도 있다.


## 입력과 출력

입력과 출력을 담당하는 것은 `io` 패키지이다. 알아둘 인터페이스는 `io.Reader`와 `io.Writer`가 있다. 책에서는 `io.Reader`를 구현하기 위한 4가지 규칙을 언급하고 있다. 근데, 나는 딱히 쓸 일은 없을 것 같다. 바로 예제나 구현해보자.

```go
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("안녕하세요."))
	fmt.Fprintln(&b, "Golang!")
	b.WriteTo(os.Stdout)
}
```

위의 코드는 다음과 같다.

```go
var b bytes.Buffer
b.Write([]byte("안녕하세요."))
```

먼저 `Buffer`에 값을 생성한 후, 문자열을 저장한다. 

```go
fmt.Fprintln(&b, "Golang!")
```

그리고 버퍼에 문자열을 결합하기 위해 `Fprintln`을 쓴다. 이 때, 첫 번째 매개변수에 `io.Writer`가 필요하다. 그래서, `bytes.Buffer`를 넘겨준다.

```go
b.WriteTo(os.Stdout)
```

그리고 해당 버퍼의 값들을 표준 출력에 출력시킨다. 이런 구조이다. 사실, `Go` 프로그래밍 할 때 숨 쉬듯 쓰고 있어서 예제의 의미를 잘 모르겠다. 하하.. 뭐 그 다음 예제로 `curl` 만드는 코드가 있는데 별의미가 없어 보인다.