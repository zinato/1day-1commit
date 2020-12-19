# 스터디 할래 - (5) 클래스

![logo](../logo.png)

> 백기선님의 온라인 스터디 "스터디 할래" 5주차 정리 문서입니다. 

## 목표

자바가 제공하는 제어문을 학습하세요.

**학습할 것**

* 클래스 정의하는 방법
* 객체 만드는 방법 (new 키워드 이해하기)
* 메소드 정의하는 방법
* 생성자 정의하는 방법
* this 키워드 이해하기

**과제 (선택)**

* int 값을 가지고 있는 이진 트리를 나타내는 Node 라는 클래스를 정의하세요.
* int value, Node left, right를 가지고 있어야 합니다.
* BinrayTree라는 클래스를 정의하고 주어진 노드를 기준으로 출력하는 bfs(Node node)와 dfs(Node node) 메소드를 구현하세요.
* DFS는 왼쪽, 루트, 오른쪽 순으로 순회하세요. 


## 클래스 정의하는 방법

클래스를 정의하는 방법은 다음과 같다. 먼저 `클래스_이름.java`라는 파일을 만들고 아래와 같은 방법으로 작성한다.

```java
public class 클래스_이름 {
    접근_제한자 타입 필드_이름;

    // ...

    접근_제한자 클래스_이름() {

    }

    // ...

    접근_제한자 타입 메소드_이름(파라미터..) {
        // ...
    }
}
```

클래스는 보통 필드, 메소드, 생성자로 이루어져 있다. 생성자는 클래스 이름으로 정의하고, 필드들을 초기화하는 역할을 한다. 

예를 들어 사람은 먼저 다음과 같은 속성이 있다.

* 나이
* 이름

이를 나타내는 `Person` 클래스는 다음과 같이 정의할 수 있다. (먼저 Person.java를 만들어둔다.)

```java
public class Person{
    int age;
    String name;
}
```

`age`는 나이를, `name`은 이름을 나타낸다. 이들을 **필드**라고 한다. 그리고 클래스에 필드/메소드/생성자를 선언할 때, "접근 제한자"라는 것을 붙인다. 접근 제한자란 외부에서 클래스의 내부 값들을 접근할 때 허용 범위를 나타낸다. 기본적으로 자바에서 제공하는 접근 제한자는 다음과 같다.

* private : 해당 클래스에서만 접근 가능하다.
* protected : 자식을 상속하는 클래스들까지 접근이 가능하다.
* public : 모든 곳에서 접근이 가능하다.
* default : 같은 "패키지"에서 접근이 가능하다.

만약 `Person`을 다음과 같이 정의했다고 해보자.

```java
public class Person{
    private int age;
    private String name;
}
```

이제 `age`, `name`은 `Person`객체만 접근이 가능하다. 이런 느낌이다.


## 객체 만드는 방법 (new 키워드 이해하기)

**다시 `private`을 제거한다.** 객체를 만들려면 `new` 연산자를 사용해야 한다. 다음과 같이 쓸 수 있다.

```java
Person p = new Person();
```

자바에서는 생성자를 선언하지 않으면, 기본 생성자를 바이트 코드로 변환하기 전에 미리 만들어준다. 만들어진 기본 생성자는 클래스 내부의 필드를 기본 값으로 초기화한다.

* 기본 타입
  * char - `\u0000`
  * byte - 0
  * short - 0
  * int - 0
  * long - 0L
  * float - 0.0f
  * double - 0.0
  * boolean - false
* 참조 타입 - null

즉, `p`의 `age`는 0, `name`은 `null`값으로 초기화된다.

```java
System.out.println(p.age);  // 0 출력
System.out.println(p.name); // null 출력
```

## 메소드 정의하는 방법

메소드는 객체의 동작이라고 말할 수 있다. `Person`은 말할 수 있어야 한다. 그냥 "안녕하세요 나는 누구이고 몇살입니다"라고 말하는 동작을 정의해보자. 다음과 같이 정의할 수 있다.

```java
public class Person{
    int age;
    String name;

    public void greeting() {
        System.out.println("안녕하세요 나는 " + name + "이고 " + age + "살입니다");    
    }
}
```

그럼 "객체를 생성한 후" `greeting` 메소드를 호출하면 된다.

```java
Person p = new Person();
p.age = 30;
p.name = "구르미"
p.greeting(); // "안녕하세요 나는 구르미이고 30살입니다." 출력
```


## this 키워드 이해하기

음 순서를 바꾸겠다. 먼저 `this`라는 것을 이해하기 위해서 다시 필드들을 `private`으로 바꾼다.

```java
public class Person{
    private int age;
    private String name;
}
```

그 후 이전에 실행했던 코드를 실행해보라.

```java
Person p = new Person();
p.age = 30;  // 컴파일 에러
p.name = "구르미";
p.greeting(); 
```

실행조차 안된다. `private` 접근 제한자가 적용되었기 때문에 말 그대로 그 "객체"만 접근이 가능하다. 그래서 컴파일 오류가 발생하는 것이다. `private`으로 접근이 막힌 필드들을 접근하기 위해서는 `this` 키워드를 사용해야 한다.

```java
public class Person{
    private int age;
    private String name;

    public void setAge(int age){
        this.age = age;
    }

    public int getAge(){
        return age;
    }

    public void setName(String name){
        this.name = name;
    }

    public String getName(){
        return this.name;
    }

    public void greeting() {
        System.out.println("안녕하세요 나는 " + name + "이고 " + age + "살입니다");    
    }
}
```

`getAge`를 보자. 

```java
return age;
```

이 때 위 구문은 "return this.age;"가 함축된 것이다. 그래서 해당 객체의 필드의 값이 반환이 된다.

`setAge`를 보자. 파라미터로 필드와 이름이 같은 `age`가 선언이 되어 있다. 그러면, 이 때 객체라고 하더라도, 파라미터 이름이 같으면 파라미터의 변수가 우선권을 가진다. 그럼 어떻게 `age` 필드를 접근해야 할까.

바로 `this` 키워드이다. 

```java
this.age = age;
```

그럼 해당 객체의 `age` 필드에 메소드 파라미터로 전달된 `age` 변수의 값이 할당된다. 이제 실행문을 다음과 같이 바꾸면 된다.

```java
Person p = new Person();
p.setAge(30);  
p.setName("구르미");
p.greeting(); // "안녕하세요 나는 구르미이고 30살입니다." 출력
```

기억하자! 객체의 내부 필드/메소드/생성자를 접근하는 것이 바로 "this" 키워드이다.


## 생성자 정의하는 방법

이제 생성자를 알아보자. 아무것도 선언이 안된어있다면, 추후에는 이렇게 코드가 변형된다.

```java
public class Person{
    private int age;
    private String name;

    public Person() {

    }

    // ...
}
```

이래서 어떤 필드도 초기화가 안되어 있는 것이다. `Person`을 다음과 같이 변경해보자.

```java
public class Person{
    private int age;
    private String name;

    public Person() {
        this.age = 15;
        this.name = "test";
    }

    // ...
}
```

그 다음 다음 구문을 실행해보자.

```java
Person p = new Person();
System.out.println(p.getAge());  // 15 출력
System.out.println(p.getName()); // test 출력
```

기본적으로 생성자는 메소드라고 볼 수 있다. 이 때 "오버로딩"이라는 것을 할 수 있다. 같은 메소드 이름이어도 파라미터가 다르면 여러 개 선언할 수 있다. 이렇게 말이다.

```java
public class Person{
    // ...

    public Person() {
    }

    public Person(int age) {
        this.age = age;
    }

    public Person(String name) {
        this.name = name;
    }

    public Person(int age, String name) {
        this.age = age;
        this.name = name;
    }

    // ...
}
```

이렇게 할 수 있다는 것이다. 이번엔 여러 생성자를 한 번에 다 써보자.

```java
Person p1 = new Person();
Person p2 = new Person(15);
Person p3 = new Person("test");
Person p4 = new Person(20, "test2");

p1.greeting(); // "안녕하세요 나는 null이고 0살입니다." 출력
p2.greeting(); // "안녕하세요 나는 null이고 15살입니다." 출력
p3.greeting(); // "안녕하세요 나는 test이고 0살입니다." 출력
p4.greeting(); // "안녕하세요 나는 test2이고 20살입니다." 출력
```

그리고 한 가지 더! 생성자에서 `this`를 조금 다르게 쓸 수 있다. 바로 `this()`로 다른 생성자를 호출하는 것인데, 코드를 다음과 같이 변경해보자.

```java
public class Person{
    // ...

    public Person() {
        this(0);
    }

    public Person(int age) {
        this(age, "test");
    }

    public Person(String name) {
        this(15, "test2")
    }

    public Person(int age, String name) {
        this.age = age;
        this.name = name;
    }

    // ...
}
```

그 후 아까 코드를 실행해보자. 한 번 아래 코드의 결과를 예상해보는 것도 좋을 것이다.


```java
Person p1 = new Person();
Person p2 = new Person(15);
Person p3 = new Person("test");
Person p4 = new Person(20, "test2");

p1.greeting(); // "안녕하세요 나는 test이고 0살입니다." 출력
p2.greeting(); // "안녕하세요 나는 test이고 15살입니다." 출력
p3.greeting(); // "안녕하세요 나는 test2이고 15살입니다." 출력
p4.greeting(); // "안녕하세요 나는 test2이고 20살입니다." 출력
```

무슨 일이 벌어진 것일까? 먼저 2번째 3번째 생성자들은, 4번째 생성자를 호출하게 된다. 두 번째 생성자는 처음에 파라미터 age에 15이 할당되고 4번째 생성자를 age=15, name="test"로 파라미터로 전달하여 객체를 초기화한다. 세번째 생성자도 동작 방식은 비슷하다.

첫 번째 생성자는 먼저 두 번째 생성자를 호출한다. 이 때 age=0으로 전달한다. 그 후 두 번째 생성자에서 age=0, name="test"를 네 번째 생성자로 전달하여 객체를 초기화한다.

## 과제
