# 1장 Go 개발 환경 설정

## 향상된 코드 포매팅  도구 goimports

`Go` 에서는 원래 포매팅 도구 `go fmt`를 지원한다. 보통 이렇게 쓴다.

```bash
go fmt hello.go
```

하지만 이 도구는 import 구문을 정리해주지 않는다. import 구문까지 정리해주는 도구 `goimports`라는 향상된 포매팅 도구가 있다. 다음과 같이 설치할 수 있다.

```bash
go install golang.org/x/tools/cmd/goimports@latest
```

현재 디렉토리 내 `.go` 파일들에 대해 포매팅하고 싶다면 다음과 같이 사용하면 된다.

```bash
# 현재 위치 확인
pwd
# 포매팅하고 싪은 go 파일들의 최상단 디렉토리
/Users/gurumee/Workspace/1day-1commit/gurumee-note/study/book/learning-go/examples

# 포매팅
goimports -l -w .
```

하지만 최신 버전 `go fmt`는 import 구문까지 잘 정리해준다. 다음 문서를 보면 더 잘 정리해주는 것을 확인할 수 있다. 

[goimports vs gofmt](https://codewithyury.com/goimports-vs-gofmt/)

## 린팅과 베팅

린팅은 코드 스타일 가이드를 잘 지켰는지 확인해주는 도구로써 `golint`를 사용한다. 다음 명령어로 설치할 수 있다.

```bash
go install golang.org/x/lint/golint@latest
```

프로젝트 전체에 린팅을 하고 싶다면 다음과 같이 사용하면 된다.

```bash
# 현재 위치 확인
pwd
# 프로젝트 루트 디렉토리
/Users/gurumee/Workspace/1day-1commit/gurumee-note/study/book/learning-go

# 린팅
golint .
```

또한, 문법적으로 유효하지만 의도한 대로 수행되지 못하는 실수들을 찾아내는 베팅 도구인 `vet`이 있다. 다음과 같이 사용할 수 있다. `go` 커맨드라인 도구에 내장되어 있어서 따로 설치는 필요 없다.

만약 다음과 같은 `go` 파일이 있다고 해보자.

learning-go/examples/ch01-vet/main.go 
```go
package main

import "fmt"

func main() {
	fmt.Printf("%d", "")
	// fmt.Printf("%d", 7)
}
```

원래는 "%d" 뒤에 인자로 숫자가 와야 한다. 하지만 의도적으로 잠재적 오류를 발생시키기 위해서 문자열을 인자로 주었다. 이제 베팅을 진행해보자.

```bash
# 현재 위치 확인
pwd
# 프로젝트 루트 디렉토리
/Users/gurumee/Workspace/1day-1commit/gurumee-note/study/book/learning-go

# 베팅
go vet examples/ch01-vet/main.go 
# command-line-arguments
examples/ch01-vet/main.go:6:2: Printf format %d has arg "" of wrong type string
```

그 다음 다음과 같이 수정해보라.

learning-go/examples/ch01-vet/main.go 
```go
package main

import "fmt"

func main() {
	// fmt.Printf("%d", "")
	fmt.Printf("%d", 7)
}
```

그 다음 베팅을 진행해보면 아무 출력이 없다. 이는 문제 없이 잘 진행된다는 의미이다.

```bash
# 현재 위치 확인
pwd
# 프로젝트 루트 디렉토리
/Users/gurumee/Workspace/1day-1commit/gurumee-note/study/book/learning-go

# 베팅
go vet examples/ch01-vet/main.go 
```

그 외 `golint`와 `vet`을 통합해주는 `golangci-lint`라는 도구가 있다. 책에서는 다음과 같이 상황에 맞게 적절한 도구를 사용할 것을 권한다.

*  프로젝트 자동화 빌드 프로세스 -> go vet
*  코드 리뷰 프로세스 -> golint
*  둘 다 익숙해지면 -> golangci-lint

## Makefiles

`make` 도구를 활용하면, 언제 어디서 누구나 실행이 가능하도록 스크립트로 자동화한다. 다음과 같이 작성할 수 있다.

learning-go/Makefile
```makefile
.DEFAULT_GOAL := build
EXAMPLE_DIR?=ch01-hello-go

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build examples/$(EXAMPLE_DIR)/main.go
.PHONY:build
```

위는 기본적으로 `build` 단계까지 실행하게 한다. `lint`, `vet`은 `fmt`에 의존하며 `build`는 `vet`에 의존하게 된다. 다음과 같이 명령어를 입력하면 해당 디렉토리의 `main.go`를 찾아 빌드하게 된다.

```bash
# 현재 위치 확인
pwd
# 프로젝트 루트 디렉토리
/Users/gurumee/Workspace/1day-1commit/gurumee-note/study/book/learning-go

# 기본 빌드: ch01-hello-go
make 
# make 출력
go fmt ./...
go vet ./...
go build examples/ch01-hello-go/main.go
# 실행 
./main
Hello, Go!

# 타겟 지정 빌드: ch01-vet
make EXAMPLE_DIR=ch01-vet
# make 출력
go fmt ./...
go vet ./...
go build examples/ch01-vet/main.go
# 실행 
./main
7%   
```