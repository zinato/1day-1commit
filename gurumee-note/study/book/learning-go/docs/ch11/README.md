# 11장 표준 라이브러리

p313. "파이썬처럼 Go도 응용 프로그램을 만들 때 필요한 많은 도구를 제공하는 '건전지 포함' 철핡을 가진다는 것이다."

## 입출력 관련 기능

p316. "7.6절 '인터페이스는 타입에 안정적인 덕 타이핑이다.'에서 논의된 바와 같이 io.Reader와 io.Writer의 구현은 종종 데코레이터 패터으로 함께 연경된다." 

> [아래 예제가 데코레이터 패턴인가?](https://refactoring.guru/design-patterns/decorator/go/example)
>
> Decorator는 Decorator라고 하는 특수 래퍼 객체 내부에 배치하여 객체에 새로운 동작을 동적으로 추가할 수 있는 구조적 패턴(Structural pattern)입니다.

p318. "이것은 사용자 지정 타입을 사용하여 더 명확하게 운영이 되었어여 했지만 설계 감시 소홀로 whence는 타입이 int가 되었다."

## 시간

p320. "Go는 일반적으로 과거에 잘 동작했던 아이디어를 체택했지만, 자체적으로 날짜 및 시간 포매팅 언어를 사용한다." -> 직관적이지가 않다. 실제로 코드를 쳐봐도 왜 이런 결과가 나오는지 이해하기 어렵다.

p322. "기본 time.Ticker는 중단할 수 없기 때문에 사소한 프로그램의 외부에 time.Tick을 사용하지 말자. 대신에 채널을 기다릴 뿐만 아니라 Ticker를 리셋하거나 중지할 수 있는 메서드를 가지는 *time.Ticker를 반환하는 time.NewTicker를 사용하자." -> Go concurrency 패턴 중에 사용 사례가 있다. Go In Action

## encoding/json

## net/http

