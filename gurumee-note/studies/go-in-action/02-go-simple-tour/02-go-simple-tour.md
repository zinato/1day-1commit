# Go 간단히 살펴보기

![대표사진](../logo.png)

> 책 "Go In Action"을 공부하면서 정리한 문서입니다.


## 목차
  - [프로그램 아키텍처](#프로그램-아키텍처)
  - [data/data.json](#datadatajson)
  - [search/feed.go](#searchfeedgo)
  - [search/match.go](#searchmatchgo)
  - [search/search.go](#searchsearchgo)
  - [search/default.go](#searchdefaultgo)
  - [matchers/rss.go](#matchersrssgo)
  - [main.go](#maingo)


## 프로그램 아키텍처

이번 장에서는 간단한 프로젝트를 진행하면서, `Go`에 대략적인 내용을 살펴본다. 애플리케이션의 아키텍처는 다음과 같다.

![아키텍처](./01.JPG)

디렉토리 폴더 구조는 다음과 같다.

```
├── ch01 ..
├── ch02
│   ├── data
│   │   └── data.json
│   ├── main.go
│   ├── matchers
│   │   └── rss.go
│   └── search
│       ├── default.go
│       ├── feed.go
│       ├── match.go
│       └── search.go
├── ....
```

자세한 것은 [여기](https://github.com/gurumee92/go-in-action)를 참고하라.

이 장을 공부하면서, 책에서는 내용을 설명하기 `main.go` 부터 필요 코드를 찾아서 올라가는 형태로 구성되어 있다. 개인적으로 칠 때 빨간색이 뜨면 짜증나기 때문에, 빨간색이 안나게끔 장을 재 구성하였다.


## data/data.json

먼저 피드를 검색할 사이트에 대한 데이터이다. 별 다른 설명은 필요 없을 듯 싶다.

go-in-action/ch02/data/data.json
```json
[
    {
        "site" : "npr",
        "link" : "http://www.npr.org/rss/rss.php?id=1001",
        "type" : "rss"
    },
    {
        "site" : "cnn",
        "link" : "http://rss.cnn.com/rss/cnn_topstories.rss",
        "type" : "rss"
    },
    {
        "site" : "foxnews",
        "link" : "http://feeds.foxnews.com/foxnews/opinion?format=xml",
        "type" : "rss"
    },
    {
        "site" : "nbcnews",
        "link" : "http://feeds.nbcnews.com/feeds/topstories",
        "type" : "rss"
    }
]
```

## search/feed.go

`data.json`의 데이터를 `Go`에서 사용할 수 있도록 만든 구조체와 그 파일을 읽어서, `Feed` 슬라이스를 만드는 코드는 `search/feed.go`에 있다.

go-in-action/ch02/search/feed.go
```go
package search

import (
	"encoding/json"
	"os"
)

const dataFile = "ch02/data/data.json"

// Feed is model of feed
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds is retrieve feeds to data/data.json
func RetrieveFeeds() ([]*Feed, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// defer는 함수 컨텍스트가 끝나는 동시에 실행된다.
	defer file.Close()
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
```

`Go`에서 구조체를 만들기 위해서는 다음과 같은 문법이 필요하다.

```
type <구조체 이름> struct {
    <필드 이름> <필드 타입>
}
```

여기서 다음 코드는 구조체의 필드가 JSON 형식으로 표현될 때의 이름을 나타낸다.

```go
type Feed struct {
    // `json:"site"`는 JSON 형식일 때 "site": "NAME 값"이 된다.
    Name string `json:"site"` 
    // ...
}
```

`RetrieveFeeds` 함수는 `ch02/data/data.json`을 읽어서 에러가 없다면, JSON 데이터들을 토대로 `Feed` 슬라이스(배열)을 만든다. 에러가 있으면, 에러를 반환한다.


## search/match.go

`search/match.go`에는 다음 코드들이 들어있다. 

- 피드를 검색한 결과를 표현하는 구조체 `Result`
- 검색을 직접적으로 수행하는 동작을 나타내는 `Matcher` 인터페이스
- 등록된 `matcher`가 `Search`를 정상적으로 수행하면, 결과를 채널에 입력하는 `Match` 함수
- `Match` 함수에서 입력된 결과를 출력하는 `Display` 함수

코드는 다음과 같다.

go-in-action/ch02/search/match.go
```go
package search

import "log"

// Result is model of result
type Result struct {
	Field   string
	Content string
}

// Matcher is interface feed, and searchTerm
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match is amtch function
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result
	}
}

// Display is fucntion display result
func Display(results chan *Result) {
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
```

여기서 `chan`이란 키워드가 "채널"을 나타낸다. 채널은 고루틴에서 데이터를 수신하고 다른 고루틴으로 송신할 수 있는 다리 역할을 해준다. 


## search/search.go

`search/search.go`는 다음 코드들이 들어 있다.

* "문자열"을 키로, 값을 `matcher`를 담는 `matchers` 맵
* `mathcers` 맵에, feedType에 따라, matcher를 등록하는 `Register` 함수
* 애플리케이션 비지니스 로직이 들어 있는 `Run`

go-in-action/ch02/search/search.go
```go
package search

import (
	"log"
	"sync"
)

// 이렇게 소문자로 변수/함수/구조체/메서드를 만들면 패키지 내부에서만 사용할 수 있다.
var matchers = make(map[string]Matcher)

// Register is register matcher
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "검색기가 이미 등록되었습니다.")
	}

	log.Println("등록 완료", feedType, " 검색기")
	matchers[feedType] = matcher
}

// Run is function exact search login
func Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatalln(err)
	}

	results := make(chan *Result)
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(feeds))

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		// Go의 익명 함수.
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	Display(results)
}
```

여기서 살펴볼 것은 이 부분이다.

```go
results := make(chan *Result)
```

이 코드는 "고루틴"을 통해 반환될 `Result` 객체를 수신할 채널을 만든다. 그리고 다음 부분을 살펴보자.

```go
go func(matcher Matcher, feed *Feed) {
    Match(matcher, feed, searchTerm, results)
    waitGroup.Done()
}(matcher, feed)
```

이 부분은 `Go`의 "익명 함수"와 "고루틴"을 살펴볼 수 있다. 우선 익명 함수이다.

```go
func(matcher Matcher, feed *Feed) {
    Match(matcher, feed, searchTerm, results)
    waitGroup.Done()
}(matcher, feed)
```

익명 함수란 이름 그대로 이름이 없는 함수이다. 함수는 보통 이렇게 짜여진다.

```
func <함수 이름> (<파라미터 이름> <파라미터 타입>, ...) (반환 타입 ...) {
    // 코드 본문
}
```

익명 함수는 여기서 이름이 빠진다. 그리고 본문 뒤에 `(matcher, feed)`가 붙는데, 이게 함수의 입력값으로 들어가서 바로 실행이 된다. 

또한 `Go`에서 비동기적인 작업을 만드는 것, "고루틴"을 만드는 것은 매우 쉽다. 바로 함수 호출 시, `go` 키워드를 붙여주는 것이다.

```
go <함수 호출>
```

즉 위의 코드는 반복문이 도는 횟수만큼 익명 함수 실행하는 고루틴들을 만든다. 그 결과를 `results` 채널에 입력하는 코드이다.


## search/default.go

`search/default.go`는 아무것도 하지 않는 `defaultMatcher`가 들어있다. 또한, `Matcher` 인터페이스를 구현하기 위한 구조체의 `Search` 메서드가 들어있다.

> 여기서 함수는 구조체에 종속되지 않는 것, 메서드는 구조체에 종속되는 것을 뜻한다. 함수는 그냥 "함수 이름"으로 부를 수 있지만, 메서드는 "<구조체 인스턴스>.메서드"로 호출할 수 있다.

go-in-action/search/default.go
```go
package search

type deafultMatcher struct{}

func init() {
	var matcher deafultMatcher
	Register("default", matcher)
}

func (m deafultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
```

여기서 `init` 함수는 패키지가 사용될 때, 초기화시키는 구문이다. 즉 `search` 패키지가 호출될 때, 자동으로 `defaultMatcher`를 `search/match.go`의 맵 `matchers`에 등록시킨다.


## matchers/rss.go

`matchers/rss.go`는 `rss`를 표현하는 타입 구조체들과, 이를 검색해서 파싱하는 `rssMatcher`와 그 메서드 `Search` 그리고 `Search` 메서드 내부에서 사용될 `rss`를 파싱하기 위한 `retrieve` 함수가 들어있다.

또한, 패키지 호출 시, `rssMatcher`를 등록할 `init` 함수가 있다.

go-in-action/ch02/matchers/rss.go
```go
package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gurumee92/go-in-action/ch02/search"
)

type (
	// item defines the fields associated with the item tag
	// in the rss document.
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image defines the fields associated with the image tag
	// in the rss document.
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel defines the fields associated with the channel tag
	// in the rss document.
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument defines the fields associated with the rss document.
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// rssMatcher implements the Matcher interface.
type rssMatcher struct{}

// init registers the matcher with the program.
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// Search looks at the document for the specified search term.
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)

	// Retrieve the data to search.
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		// Check the title for the search term.
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		// If we found a match save the result.
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		// Check the description for the search term.
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		// If we found a match save the result.
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}

// retrieve performs a HTTP Get request for the rss feed and decodes the results.
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed uri provided")
	}

	// Retrieve the rss feed document from the web.
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// Close the response once we return from the function.
	defer resp.Body.Close()

	// Check the status code for a 200 so we know we have received a
	// proper response.
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Http Response Error %d\n", resp.StatusCode)
	}

	// Decode the rss feed document into our struct type.
	// We don't need to check for errors, the caller can do this.
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
```

## main.go

그리고 마지막으로, 프로그램을 시킬 `main.go`가 있다. 

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

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("corona")
}
```

`Go`에서는 실행 가능한 바이너리 파일이 되려면, `package main`과 `main` 함수 선언이 필수적이다. 

그리고, 표준 패키지 `log`, `os`를 임포트한다. 또한, 외부 패키지를 가져오기 위해서 그 주소를 입력하면 된다.

```
_ "github.com/gurumee92/go-in-action/ch02/matchers"
```

위의 임포트 구문은 `matchers`를 사용하지는 않지만, `matchers` 패키지에서 초기화시키고 싶은 것이 있을 때 명시적으로 패키지 이름을 주는 것이다.

> Go에서는 쓰지 않는 변수, 패키지 이름이 있을 경우, 컴파일이 되지 않는다. 쓰지 않는 이름이 있을 때, 이렇게 "_" 써서 넘어갈 수 있다.

