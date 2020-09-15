# Go의 타입 시스템

![대표사진](../logo.png)

> 책 "Go In Action"을 공부하면서 정리한 문서입니다.


## 목차
  - [기본 타입과 참조 타입](#기본-타입과-참조-타입)
  - [사용자 정의 타입](#사용자-정의-타입)
  - [구조체와 메서드](#구조체와-메서드)
  - [구조체와 인터페이스](#구조체와-인터페이스)
  - [구조체와 타입 임베딩](#구조체와-타입-임베딩)


## 기본 타입과 참조 타입

`Go`의 타입은 크게 2가지로 나눌 수 있다. 기본 타입과 참조 타입이다. 이 둘은 또 이렇게 분류할 수 있겠다.

* 기본 타입 : int, float, bool, rune 등의 값 표현 타입, 포인터, 배열, string
* 참조 타입 : 구조체, 인터페이스, 슬라이스, 맵, 채널, 함수

이 둘의 차이점은 참조 타입은 기본적으로 헤더 값을 가지고 있다. 기술적으로 보면 문자열을 표현하는 `string` 역시 참조 타입으로 볼 수 있다. 그러나, 동작하는 것은 기본 타입에 가까워서 기본 타입으로 분류하였다.


## 사용자 정의 타입

기본 타입 외, 타입을 정의하는 방법은 크게 2가지이다. 바로 `구조체`와 `타입 재정의`이다. 타입 재정의는 있는 타입을 다시 정의하는 것이다. 다음처럼 말이다.

```go
type MyInt int
```

이는 `int` 타입을 `MyInt`로 타입을 재정의한 것이다. 실제로 이런 식으로 초기화할 수 있다.

```go
var i1, i2 MyInt
i1 = 5
i2 = MyInt(5)
```

그런데 원래 타입을 다시 할당하면 컴파일 에러가 뜬다.

```go
var i3 int = 5
i1 = i3 // 이러면 컴파일 에러
```

왜냐하면 `MyInt` 타입에 `int` 값을 할당했기 때문이다. 둘은 엄연히 다른 타입이기 때문에 타입 캐스팅을 하지 않는 한 이런 식의 할당은 불가능하다. 또 다른 방법은 구조체이다. 구조체는 여러 타입을 한 타입으로 묶을 수 있다. 다음처럼 말이다.

```go
type Point struct {
	X int
	Y int
}
```

여기서, `X`, `Y`를 구조체의 "멤버 필드"라고 부른다. 내부의 속성을 가리킨다. 위의 예제는 `int` 타입 두 개를 묶어주었지만, 여러 타입은 물론 다른 구조체, 인터페이스까지 묶을 수 있다. 초기화는 다음의 2가지 방식이 있다.

```go
p1 := Point{
    X: 1,
    Y: 2,
}
p2 := Point{4, 1}
```

첫 번째 방식은 멤버 필드를 명시적으로 선언해서 초기화한다. 두 번째 방식은 선언한 순서대로 초기화된다. 첫 번째 방식에선, `X`, `Y`의 순서가 바뀌어도 구조체의 값은 같지만 두 번째 방식에선 `X`, `Y`의 값이 바뀐다.


## 구조체와 메서드

`멤버 필드`는 구조체의 속성을 나타낸다고 하였다. `메서드`는 구조체의 행위를 나타낸다. 구조체에서 호출하는 함수라고 생각하면 편하다. 만약 다음의 구조체가 선언되었다고 보자.

```go
type Person struct {
	Name string
	Mail string
}
```

그리고 메서드는 다음처럼 선언 및 정의할 수 있다.

```go
func (p Person) GetMail() string {
	return p.Mail
}

func (p *Person) SetMail(mail string) {
	p.Mail = mail
}
```

코드 구조는 다음과 같다.

```
func ((*)구조체) 함수 이름(함수 파라미터) { 정의 }
```

여기서 `func` 뒤의 `( (*)구조체 )` 자리는, "수신자"라고 한다. 이렇게 수진자를 선언하면, 구조체의 메서드를 선언한 것이다. 그렇다면 `*`가 붙은 것과 안 붙은 것의 차이는 무엇일까? 이건 함수랑 같다. 포인터를 함수 파라미터로 전달하면 값이 바뀌었듯이, 포인터를 메서드 수신자로 전달하면, 구조체 내의 필드를 수정할 수 있다. 

```go
p := Person{
    Name: "Gurumee",
    Mail: "gurumee@example.com",
}
p.SetMail("gara@example.com")
```

실제 이 경우 `p`의 멤버 필드 `Mail`의 값이 `SetMail` 메서드 호출 후 바뀐 것을 확인할 수 있다. 그러나 `SetMail`에서 `*`를 빼면, 메서드 호출 이후에도 값이 바뀌지 않는다.


## 구조체와 인터페이스

`Go`는 인터페이스를 통한 다형성을 지원한다. 인터페이스 중심의 OOP를 지원한다. 인터페이스는 명세라고 보통 알고 있지만 `Go`에서는 인터페이스 1개당 1개의 메서드를 선언하는 것을 권장하고 있다. 인터페이스 선언은 다음과 같다.

```go
type Notifier interface {
	Notify()
}
```

메서드 시그니처만 선언 후, 정의는 따로 하지 않는다. 인터페이스를 구현하려면 어떻게 해야할까? 그냥 구조체가 저 메서드를 구현하면 된다.

```go
type Person struct {
	Name string
	Mail string
}

// ...

func (p *Person) Notify() {
	fmt.Println(p.Name, p.Mail, "로 전송합니다.")
}
```

뭐가 좋은 것일까? 우선 이런 식의 코드를 작성할 수 있다.

```go
p := Person{
    Name: "Gurumee",
    Mail: "gurumee@example.com",
}
p.Notify()

var n Notifier = &p
n.Notify()
```

여기서 수신자가 `*Person`이기 때문에, 주소를 넘겨주어야 한다. 이제 다음의 함수를 구현한다고 가정해보자.

```go
func SendNotify(n Notifier) {
	n.Notify()
}
```

그러면, 이렇게 호출할 수 있을 것이다.

```go
SendNotify(&p)
```

이렇게 하면 인터페이스의 장점이 보이지 않는다. 근데, 관리자 타입을 구현해야 한다고 해보자. 이 관리자 타입도 `Notifier` 인터페이스를 구현해야 한다고 해보자.

```go
type Admin struct {
	Name  string
	Mail  string
	Level string
}

func (a *Admin) Notify() {
	fmt.Println(a.Name, a.Mail, a.Level, "로 전송합니다.[관리자]")
}
```

그럼 아래 코드처럼 할당 및 함수를 호출할 수 있을 것이다.

```go
a := Admin{
    Name:  "ADMIN",
    Mail:  "ADMIN@example.com",
    Level: "SUPER",
}
SendNotify(&a)
```

구조체 기반으로 함수 등을 작성하면, 그만큼 유연성이 떨어진다. 만약 `SendNotify`와 수 천, 수 만개의 함수들이 `Person` 기반으로 작성되었다고 해보자. 이 때 `Admin` 기반으로 다 바꿔야 한다고 해보자. 그러면, 수 천, 수 만번의 함수를 재작성해야 한다. 그렇지만 인터페이스 기반으로 작성하면, 구현하는 구조체와 메서드만 작성하면 된다. 코드의 유연성이 생기게 된다.


## 구조체와 타입 임베딩

`Go`에서는 타입 임베딩을 지원한다. `Admin`을 다음처럼 바꿔보자. 그리고 `Admin`이 구현하는 `Notify` 메서드를 제거해보자.

```go
type Admin struct {
	Person
	Level string
}

// func (a *Admin) Notify() {
// 	fmt.Println(a.Name, a.Mail, a.Level, "로 전송합니다.[관리자]")
// }
```

이 경우 인터페이스 `Notifier`의 `Notify`를 `Admin`이 구현하지 않는다. 그러면 `SendNotify`가 실행되지 않지 않을까? 그러나 놀랍게도 실행된다. `타입 임베딩`은 위와 같이 멤버 필드의 이름 없이 타입만 구조체 내에 선언하게 되면, 내부적으로 그 구조체가 구현하는 필드와 메서드를 모두 가질 수 있다. 

그리고 해당 인터페이스를 구현하는 것을 임베딩하고 있을 경우, 그 인터페이스로 사용되야 할 때 자동으로 그 인터페이스를 구현하는 내부 구조체가 승격하여 외부 구조체를 대신한다. 즉, `SendNotify`가 호출될 때 `Admin`은 내부 구조체 `Person`으로 취급된다는 것이다. 다른 말로 호출 될 때 `Person`이 승격된다라고 말할 수 있다. 하지만 다시 `Admin`이 `Notify` 메서드를 구현할 경우 다시 승격하지는 않는다. 

`Admin`구조체의 `Notify` 주석을 해제해고 그 결과를 확인해보자.

주석 해제 전
```
Gurumee gurumee@example.com 로 전송합니다.
Gurumee gurumee@example.com 로 전송합니다.
```

주석 해제 후
```
Gurumee gurumee@example.com 로 전송합니다.
Gurumee gurumee@example.com SUPER 로 전송합니다.[관리자]
```


