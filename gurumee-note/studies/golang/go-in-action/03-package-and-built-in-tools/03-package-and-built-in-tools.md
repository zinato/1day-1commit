# 패키징과 내장 도구들

![대표사진](../logo.png)

> 책 "Go In Action"을 공부하면서 정리한 문서입니다.


## 목차
  - [패키지](#패키지)
  - [main 패키지](#main-패키지)
  - [import](#import)
  - [init](#init)
  - [Go 내장 도구의 활용](#go-내장-도구의-활용)
  - [다른 Go 개발자와 협업하기](#다른-go-개발자와-협업하기)
  - [의존성 관리](#의존성-관리)


## 패키지

`Go`의 패키지는 "기능"을 "의미"적으로 분류하여, 각기 다른 패키지로 분리하는 것이다. `http` 패키지를 구성하는 하위 패키지는 다음과 같다.

```
net/http/
    cgi/
    cookiejar/
        testdata/
    fcgi/
    httptest/
    httputil/
    pprof/
    testdata/
```

이 패키지들은 모두 여러 `go` 파일들을 저장하고 있다. 모든 `go` 파일들은 주석을 제외한 최 상단에 자신이 속한 `package`를 선언해야 한다. 하나의 디렉토리는 하나의 패키지만을 표현할 수 있다. 해당 패키지의 코드들은 모두 같은 디렉토리에 위치해야 한다.

패키지 이름 규칙은 간단하다. 

* 패키지가 저장되는 디렉토리의 이름을 따른다.
* 짧고 간결하고, 소문자로만 구성한다.


## main 패키지

패키지 중에서도 제일 특별한 패키지는 `main` 패키지이다. `main` 패키지 안에 `main` 함수를 작성하지 않으면, 실행 파일이 만들어지지도 않는다. 실행 파일을 만들고 싶다면 반드시 `main` 패키지 선언 후, `main` 함수를 작성하자. (단, 반드시 "main.go"일 필요는 없다.)


## import

`import` 구문은 패키지를 가져온다. `go`가 설치된 디렉토리 밑에 하위 패키지들을 가져오기도 하고, `go mod`를 이용해서 만들어진 `vendor` 디렉토리 밑에 하위 패키지들을 가져오기도 한다. 책이 만들어진 시점에는 언어가 공식적으로 지원하는 의존성 관리 도구인 `go mod`가 없었다. 그래서 획일화되지 않았는데, 현재는 모두 `go mod`를 따르는 추세이다.

또한, `Go`는 `bitbucket`, `Github`, `Gitlab` 등의 외부 깃 서버에서 패키지를 가져올 수 있다. 이것 역시 `go mod`로 벤더링 후, `go get`을 하면 외부 패키지도 여기에 저장되기에 로컬에 패키지를 다운받았다고 해서, 코드의 주소를 바꿀 필요는 없다.

또한, 패키지 하위 패키지들이 이름이 겹칠 수가 있다. 이럴 때는, 패키지의 이름을 붙여서 `import`가 가능하다. 자신이 만든 "fmt" 패키지가 있다고 가정하고 이 둘을 `import`하는 상황이라고 가정해보자. 그럼 이런 식으로 `import`가 가능하다.

```go
package main

import (
    "fmt"
    myfmt "mylib/fmt"
)

func main() {
    fmt.Println("표준 라이브러리")
    myfmt.Println("나의 입/출력 라이브러리")
}
```


## init

지난 장에서 봤던 구문이다.

go-in-action/ch02/main.go
```go
package main

import (
	// 표준 패키지, log, os를 가져옴
	"log"
	"os"

	// 현재 모듈 matchers, search 가져옴.
	_ "github.com/gurumee92/go-in-action/ch02/matchers"
	"github.com/gurumee92/go-in-action/ch02/search"
)

// ...
```

`main.go`에서 `matchers` 패키지를 직접적으로 사용하는 코드는 없지만, `matchers` 패키지 하위에 `rss.go`의 `rssMatcher`를 등록해주는 작업이 필요하다. 이렇게 패키지를 "_"로 명명할 수 있다. 그리고 `matchers/rss.go`를 보면 `init` 함수에서 그 작업을 하는 것을 확인할 수 있다.

go-in-action/ch02/matchers/rss.go
```go
package matchers

// ...

func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

//...
```

보통 데이터베이스를 연결하는 작업 등에 많이 쓰인다.


## Go 내장 도구의 활용

`Go`는 수 많은 내장 도구들을 지원한다. 터미널에 "go"를 입력해보자.

```bash
$ go
Go is a tool for managing Go source code.

Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildmode   build modes
        c           calling between Go and C
        cache       build and test caching
        environment environment variables
        filetype    file types
        go.mod      the go.mod file
        gopath      GOPATH environment variable
        gopath-get  legacy GOPATH go get
        goproxy     module proxy protocol
        importpath  import path syntax
        modules     modules, module versions, and more
        module-get  module-aware go get
        module-auth module authentication using go.sum
        module-private module configuration for non-public modules
        packages    package lists and patterns
        testflag    testing flags
        testfunc    testing functions

Use "go help <topic>" for more information about that topic.
```

가장 많이 쓰이는 것은 `build`, `run`, `fmt`, `get`, `test`, `vet`, `mod` 정도이다. 간단히 살펴보면 다음과 같다.

* build - 프로젝트를 실행 가능한 바이너리 파일로 빌드한다. 보통은 main 패키지와 main 함수가 있는 파일의 이름을 따른다. (명명할 수 있음)
* run - build 후 실행 파일 실행한 뒤, 그 파일을 삭제한다. 인터프리터 언어처럼 사용할 수 있다. (빠른 컴파일 속도 때문에)
* test - `test`가 붙은 `.go`파일들의 테스트 코드를 실행한다. (테스트 함수의 이름은 "Test"가 들어가 있다.)
* fmt - Go 코드를 정리해준다.
* vet - Go 코드에서 일반적으로 발생할 수 있는 에러를 잡아준다.
* get - 외부 패키지를 다운로드 한다.
* mod - 의존성 관리를 해준다. `go mod vendor`를 하면 필요 패키지들을 `vendor`라는 디렉토리에 저장한다.


## 다른 Go 개발자와 협업하기

소스 코드를 공유하고 싶다면 몇 가지 규칙을 지켜야 한다.

1. 패키지는 반드시 저장소의 루트에 저장해야 한다
2. 패키지의 크기는 작게 유지하자
3. 코드에 `go fmt` 명령을 실행하자
4. 소스 코드를 문서화하자

나한테는 별로 해당 사항이 없어 보인다. 내 소스 코드는 `Github`이 공유해줄거야..


## 의존성 관리

책에서는 4가지 `godep`, `vendor`, `gopkg.in`, `gb`를 소개하는데 다 필요 없다. 이제 `go mod`로 이 모든 것을 해결할 수 있다. (`gb`가 옮겨졌거나, 영향을 받아 만들어진듯 하다.)