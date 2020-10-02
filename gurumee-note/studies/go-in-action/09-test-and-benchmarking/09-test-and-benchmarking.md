# 테스트와 벤치마킹

![대표사진](../logo.png)

> 책 "Go In Action"을 공부하면서 정리한 문서입니다.


## 기본 단위 테스트

`단위 테스트`란 프로그램/패키지의 일부 코드를 테스트하는 함수를 말한다. 먼저 기본적인 단위 테스트는 다음과 같다.

> 주요한 점은 go 파일을 만들때 끝에 "_test"를 붙여야 한다. 예를 들면 example_test.go 이런 식으로 말이다.

```go
package ch09

import (
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	const checkMark = "\u2713"
	const ballotX = "\u2717"
	url := "https://api.hnpwa.com/v0/news/1.json"
	statusCode := 200

	t.Log("컨텐츠 다운로드 시작")
	{
		t.Logf("\tURL \"%s\" check status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)

			if err != nil {
				t.Fatal("\t\t HTTP GET Check", ballotX, err)
			}

			t.Log("\t\t HTTP GET Check", checkMark)
			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\t Status Code Check \"%d\": \"%v\"", statusCode, checkMark)
			} else {
				t.Errorf("\t\t Status Code Check \"%d\": \"%v\" \"%v\"", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
```

터미널에 다음을 입력하면 된다. 이 명령은 프로젝트를 `go mod init`으로 초기화한 경우만 먹힐 것이다.

```bash
$ go test -v ./ch09
=== RUN   TestDownload
    TestDownload: example_09_01_test.go:15: 컨텐츠 다운로드 시작
    TestDownload: example_09_01_test.go:17:     URL "https://api.hnpwa.com/v0/news/1.json" check status code "200"
    TestDownload: example_09_01_test.go:25:              HTTP GET Check ✓
    TestDownload: example_09_01_test.go:29:              Status Code Check "200": "✓"
--- PASS: TestDownload (1.83s)
PASS
ok      command-line-arguments  2.476s
```

`go test`는 "*_test.go" 파일을 테스트 파일로 간주한다. 따라서, 위 명령어를 입력했을 때 "ch09" 패키지 밑에 모든 테스트 파일에서 테스트 함수들을 실행한다. 

```go
func TestDownload(t *testing.T) {
	// ...
}
```

또한, 테스트 함수는 반드시 위의 코드처럼 `Test*`로 시작해야 하며, 매개 변수로 `*testing.T`를 받는다. 리턴 값 역시 없어야 한다. `단위 테스트`는 해당 테스트가 어떤 이유에서 필요한지 서술해야 한다. 그리고 어떤 결과가 나오는지 알려주어야 한다. 뭐 대부분은 아니지만, 필자는 보통 다음의 구조로 단위 테스트를 한다.

1. 테스트를 위한 값 설정 given
2. 테스트하기 위한 코드 호출 when
3. 결과 확인 then

```go
// given
url := "https://api.hnpwa.com/v0/news/1.json"
statusCode := 200
```

이 부분이 바로 테스트를 위한 매개변수이다. `url`을 호출했을 때, 상태 코드 200이 나오느냐를 확인하기 위해 주어진 값들이다.

```go
// when
resp, err := http.Get(url)

if err != nil {
	t.Fatal("\t\t HTTP GET Check", ballotX, err)
}

// ...
defer resp.Body.Close()
```

실제적인 코드 호출이다. `url`을 호출한 것을 테스트하기 위한 호출이다.

```go
// then
if resp.StatusCode == statusCode {
	t.Logf("\t\t Status Code Check \"%d\": \"%v\"", statusCode, checkMark)
} else {
	t.Errorf("\t\t Status Code Check \"%d\": \"%v\" \"%v\"", statusCode, ballotX, resp.StatusCode)
}
```

상태 코드가 200이 나면, 체크 표시를 아니면 엑스 표시를 해서 로깅하게 한다. 또한, `t.Errorf`가 있기 때문에 실패하면, 테스트 실패가 떨어질 것이다. 실제 `https://api.hnpwa.com/v0/news/100.json` 을 넣고 실행할 때는, 다음의 결과를 확인할 수 있다.

```
=== RUN   TestDownload
    TestDownload: example_09_01_test.go:16: 컨텐츠 다운로드 시작
    TestDownload: example_09_01_test.go:17:     URL "https://api.hnpwa.com/v0/news/100.json" check status code "200"
    TestDownload: example_09_01_test.go:25:              HTTP GET Check ✓
    TestDownload: example_09_01_test.go:32:              Status Code Check "200": "✗" "500"
--- FAIL: TestDownload (4.32s)
FAIL
FAIL    github.com/gurumee92/go-in-action/ch09  4.628s
FAIL
```

사실, 테스트를 하려면 여러 매개 변수로 맞는지 확인을 해야 한다. 이를 `테이블 테스트`라 한다. 이번엔 여러 URL을 보내서 상태 코드 결과를 확인해보자. 코드르 다음과 같이 변경한다.

```go
package ch09

import (
	"net/http"
	"testing"
)

func TestDownload(t *testing.T) {
	// given
	const checkMark = "\u2713"
	const ballotX = "\u2717"
	params := []struct {
		url        string
		statusCode int
	}{
		{
			url:        "https://api.hnpwa.com/v0/news/1.json",
			statusCode: 200,
		},
		{
			url:        "https://api.hnpwa.com/v1/news/1.json",
			statusCode: 404,
		},
		{
			url:        "https://api.hnpwa.com/v0/news/100.json",
			statusCode: 500,
		},
	}

	t.Log("컨텐츠 다운로드 시작")
	for _, param := range params {
		t.Logf("\tURL \"%s\" check status code \"%d\"", param.url, param.statusCode)

		resp, err := http.Get(param.url)

		if err != nil {
			t.Fatal("\t\t HTTP GET Check", ballotX, err)
		}

		t.Log("\t\t HTTP GET Check", checkMark)
		defer resp.Body.Close()

		// then
		if resp.StatusCode == param.statusCode {
			t.Logf("\t\t Status Code Check \"%d\": \"%v\"", param.statusCode, checkMark)
		} else {
			t.Errorf("\t\t Status Code Check \"%d\": \"%v\" \"%v\"", param.statusCode, ballotX, resp.StatusCode)
		}
	}
}
```

이제 테스트 코드를 실행해보자.

```bash
$ go test -v ./ch09
=== RUN   TestDownload
    TestDownload: example_09_01_test.go:31: 컨텐츠 다운로드 시작
    TestDownload: example_09_01_test.go:33:     URL "https://api.hnpwa.com/v0/news/1.json" check status code "200"
    TestDownload: example_09_01_test.go:41:              HTTP GET Check ✓
    TestDownload: example_09_01_test.go:46:              Status Code Check "200": "✓"
    TestDownload: example_09_01_test.go:33:     URL "https://api.hnpwa.com/v1/news/1.json" check status code "404"
    TestDownload: example_09_01_test.go:41:              HTTP GET Check ✓
    TestDownload: example_09_01_test.go:46:              Status Code Check "404": "✓"
    TestDownload: example_09_01_test.go:33:     URL "https://api.hnpwa.com/v0/news/100.json" check status code "500"
    TestDownload: example_09_01_test.go:41:              HTTP GET Check ✓
    TestDownload: example_09_01_test.go:46:              Status Code Check "500": "✓"
--- PASS: TestDownload (4.92s)
PASS
ok      github.com/gurumee92/go-in-action/ch09  5.431s
```

각 URL에 맞는 상태코드가 떨어지는 것을 확인할 수 있다.


## 목 테스트

목 테스트는 단위 테스트할 수 있게, 가짜 결과를 만들어내는 목 객체를 만들어서 테스트하는 방법이다. 현재 위의 테스트 코드가 실패할 경우는 언제일까? 바로 인터넷이 막혀있을 때다. 뭐 보통 이럴 땐 없겠지만, 실제 개발에서는 다른 컴포넌트가 개발이 완료되지 않았을 때 내가 그 컴포넌트를 이용하는 또 다른 컴포넌트를 개발할 때가 있다. 이 경우에 쓸 수 있다. 해당 컴포넌트가 결과값이 이럴 것이다만 형태를 갖춰놓았다면, 이 결과를 만들어내는 가짜 객체를 생성하면 된다.

바로 코드를 살펴보자.

```go
package ch09

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var json = `
	[
		{
			id: 24661126,
			title: "Flatpak: A security nightmare – two years later",
			points: 90,
			user: "krimeo",
			time: 1601634959,
			time_ago: "an hour ago",
			comments_count: 50,
			type: "link",
			url: "https://www.flatkill.org/2020/",
			domain: "flatkill.org"
		},
	]
`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, json)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestMockDownload(t *testing.T) {
	// given
	const checkMark = "\u2713"
	const ballotX = "\u2717"
	server := mockServer()
	statusCode := 200

	t.Log("컨텐츠 다운로드 시작")
	t.Logf("\tURL \"%s\" check status code \"%d\"", server.URL, statusCode)
	// when
	resp, err := http.Get(server.URL)

	if err != nil {
		t.Fatal("\t\t HTTP GET Check", ballotX, err)
	}

	t.Log("\t\t HTTP GET Check", checkMark)
	defer resp.Body.Close()

	// then
	if resp.StatusCode == statusCode {
		t.Logf("\t\t Status Code Check \"%d\": \"%v\"", statusCode, checkMark)
	} else {
		t.Errorf("\t\t Status Code Check \"%d\": \"%v\" \"%v\"", statusCode, ballotX, resp.StatusCode)
	}
}
```

아까와 코드는 거의 동일하다. 대신 `mockServer` 함수로 만든 `server`의 URL로 바꿔쓰고 있다. 이제 `json`과 `mockServer` 함수 부분을 보자.

```go
var json = `
	[
		{
			id: 24661126,
			title: "Flatpak: A security nightmare – two years later",
			points: 90,
			user: "krimeo",
			time: 1601634959,
			time_ago: "an hour ago",
			comments_count: 50,
			type: "link",
			url: "https://www.flatkill.org/2020/",
			domain: "flatkill.org"
		},
	]
`

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, json)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}
```

`mockServer`는 HTTP 요청 시, 서버의 응답을 대신하는 객체를 만들어준다. 이 서버는 요청 마다, `json`을 결과로 내보내준다. 그래서 해당 테스트 코드를 돌려보면 다음의 결과를 얻을 수 있다.

```bash
$ go test -v ./ch09/example_09_02_test.go
=== RUN   TestMockDownload
    TestMockDownload: example_09_02_test.go:43: 컨텐츠 다운로드 시작
    TestMockDownload: example_09_02_test.go:44:         URL "http://127.0.0.1:51568" check status code "200"
    TestMockDownload: example_09_02_test.go:52:                  HTTP GET Check ✓
    TestMockDownload: example_09_02_test.go:57:                  Status Code Check "200": "✓"
--- PASS: TestMockDownload (0.00s)
PASS
ok      command-line-arguments  0.491s
```

목 객체라 URL이 마음대로 잡힌 것을 확인할 수 있다. `목 테스트`를 이용하면, 개발되지 않은 컴포넌트라도 내가 원하는 컴포넌트의 유닛테스트가 가능하다. 그러나 단점도 존재한다. 만약, 해당 컴포넌트의 결과가 합의되지 않았거나 중간에 변경된 경우 이 테스트는 쓸모 없는 테스트가 되버린다. 

그래서 `통합 테스트`라는 것이 존재한다. `단위 테스트`들을 전체적으로 연결해서 서비스가 동작하는지 테스트하는 것이다. 이 때 `목 테스트`를 쓰는 `단위 테스트`가 있다면, 목 객체를 실제 객체로 바꿔주어야 한다.

> 책에서는 엔드포인트 테스트와 문서화를 다루고 있다. 나와는 별 상관이 없어 보여서 생략한다.


## 벤치 마킹



