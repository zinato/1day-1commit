# 일반적인 프로그래밍 원칙
===

![대표사진](../intro.png)

> 책 "이펙티브 자바 3판"을 공부하면서 정리한 문서입니다.


## 지역변수의 범위를 최소화하라

책에서는 다음 상황을 권장한다.

1) 가장 처음 쓸 때 선언하기 = 선언과 동시에 초기화한다.
2) `while`보다 `for`가 낫다.
3) 메서드를 최대한 작게, 한 가지 기능에 집중하라.


## 전통적인 for문보다는 for-each문을 사용하라

전통적인 `for`보단, `for-each`를 사용한다.

```java
for (Element e: elements) {
    // ...
}
```

`for-each`를 쓰기 어려운 상황이 있다.

1) 파괴적인 필터링 - 순회하면서, 선택된 원소를 제거해야 할 때, -> 이 경우는 `removeIf`를 사용해서, 명시적 순회를 피할 수 있다.
2) 변형 - 컬렉션의 원소 혹은 전체를 교체해야 할 때, 인덱스로 접근해야 한다.
3) 병렬 반복 - 여러 컬렉션을 병렬로 순회해야 할 때


## 라이브러리를 익히고 사용하라

이번 아이템은 "**바퀴를 다시 발명하지 말자**"이다. 책에서는 표준 라이브러리를 익히면 좋은 점 5가지를 다음과 같이 소개하고 있다.

1) 코드를 작성한 전문가의 지식과 자신보다 앞서 사용한 다른 프로그래머들의 경험을 활용할 수 있다.
2) 비지니스 로직과 크게 관련 없는 문제를 해결하느라 시간을 낭비하지 않아도 된다.
3) 따로 노력하지 않아도 성능이 지속해서 개선된다.
4) 기능이 점점 많아진다.
5) 많은 사람에게 낯익은 코드가 된다. -> 유지보수하기 좋고 가독성이 좋은 코드가 된다.


## 정확한 답이 필요하다면 float와 double은 피하라

`float`, `double`은 근사치로 계산하도록 설계되었다. "정확"과는 거리가 있다는 것이다. 다음 코드를 보자.

```java
@Test
@DisplayName("실수로 테스트")
public void test01() {
    double funds = 1.00;
    int itemsBought = 0;

    for (double price = 0.10; funds >= price; price += 0.10) {
        funds -= price;
        itemsBought++;
    }

    assertNotEquals(4, itemsBought);
    assertNotEquals(0, funds);
}
```

원래 의도라면, 4개를 구입하고 0원이 남아야 하지만, 실제로는 3개, 0.399999원이 남아있다. 금융 계산처럼 정확한 계산이 필요할 땐, `BigDecimal`, `int`, `long`을 사용해라. 먼저 위의 코드를 `BigDecimal`로 바꾼 코드이다.

```java
@Test
@DisplayName("BigDecimal로 테스트")
public void test02() {
    final BigDecimal TEN_CENTS = new BigDecimal(".10");
    BigDecimal funds = new BigDecimal("1.00");
    int itemsBought = 0;

    for (BigDecimal price = TEN_CENTS; funds.compareTo(price) >= 0; price = price.add(TEN_CENTS)) {
        funds = funds.subtract(price);
        itemsBought++;
    }

    assertEquals(4, itemsBought);
    assertEquals(0, funds.intValue());
}
```

정확한 결과가 나온다. 그러나 `BigDecimal`의 단점은 다음 2가지가 있다.

1) 기본 타입에 비해 사용하기 어렵다.
2) 느리다.

이번엔 기본 타입으로 계산하기 위해서 단위를 달러에서 센트로 바꾼 코드이다.

```java
@Test
@DisplayName("정수 테스트")
public void test03() {
    int funds = 100;
    int itemsBought = 0;

    for (int price = 10; funds >= price; price += 10) {
        funds -= price;
        itemsBought++;
    }

    assertEquals(4, itemsBought);
    assertEquals(0, funds);
}
```

역시 테스트 결과는 정상적으로 나온다. 그러나 기본 타입은 거대한 실수를 표현할 때 계산이 쪼금 복잡하다. 책에서는 아홉자리는 `int`, 18자리는 `long` 그 이후 자릿 수는 `BigDecimal`을 쓸 것을 권고하고 있다.


## 박싱된 기본 타입보다는 기본 타입을 사용하라

자바의 기본 타입은 크게 2가지로 분류할 수 있다.

* int, boolean, double 등의 기본 타입
* Integer, Boolean, Double 등의 기본 타입의 래퍼 클래스

이 둘의 큰 차이점은 크게 3가지다.

1) 기본 타입은 값만, 래퍼 클래스는 값과 식별성을 갖는다.
2) 기본 타입은 언제나 유효한 값을, 래퍼 클래스는 null을 가질 수 있다.
3) 메모리 효율성은 기본 타입이 래퍼 클래스보다 좋다.

첫 번째 차이점은 조금 중요하다. 식별성을 갖는다는 것은 `==` 연산자로 동일한지 확신할 수 없다. 다음 테스트 코드의 결과를 예측해보자.

```java
@Test
@DisplayName("Wrapper 클래스 식별성 테스트")
public void test01() {
    Integer i1 = new Integer(42);
    Integer i2 = new Integer(42);
    assertTrue(i1 == i2);
}
```

이 테스트는 통과할 것인가? 예측과는 다르게 테스트는 실패한다. 왜냐하면 두 객체 `i1`, `i2`는 서로 다른 식별성을 가지고 있기 때문이다. 이럴 땐, `i1.equals(i2)`로 하거나, 기본 타입을 사용해야 올바르게 코드가 동작된다. 

박싱 타입을 쓸 때, 위의 문제점과 포함하여 다음의 3가지 단점이 존재한다.

1) 개발자가 작성한 `==` 연산자가 오동작 할 위험이 있다. (박싱 타입의 동일성 체크는 `equals` 메소드를 쓰자)
2) 널 포인터 예외가 발생할 수 있다.
3) 성능이 저하된다.

따라서 어지간하면, 기본 타입은 박싱 타입이 아닌 그냥 기본 타입을 쓰도록 하자.


## 다른 타입이 적절하다면 문자열 사용을 피하라


## 문자열 연결은 느리니 주의하라


## 객체는 인터페이스를 사용해 참조하라


## 리플렉션보다는 인터페이스를 사용하라


## 네이티브 메서드는 신중히 사용하라


## 최적화는 신중히 하라


## 일반적으로 통용되는 명명 규칙을 따르라


