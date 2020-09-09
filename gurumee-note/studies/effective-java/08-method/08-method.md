메서드
===

![대표사진](../intro.png)

> 책 "이펙티브 자바 3판"을 공부하면서 정리한 문서입니다.

## 목차
  - [매개변수가 유효한지 검사하라](#매개변수가-유효한지-검사하라)
  - [적시에 방어적 복사본을 만들라](#적시에-방어적-복사본을-만들라)
  - [메서드 시그니처를 신중히 설계하라](#메서드-시그니처를-신중히-설계하라)
  - [다중정의는 신중히 사용하라](#다중정의는-신중히-사용하라)
  - [가변인수는 신중히 사용하라](#가변인수는-신중히-사용하라)
  - [null이 아닌 빈 컬렉션이나 배열을 반환하라](#null이-아닌-빈-컬렉션이나-배열을-반환하라)
  - [옵셔널 반환은 신중히 하라](#옵셔널-반환은-신중히-하라)
  - [공개된 API 요소에는 항상 문서화 주석을 작성하라](#공개된-api-요소에는-항상-문서화-주석을-작성하라)


## 매개변수가 유효한지 검사하라

"오류는 가능한 빨리 잡아야 한다" - 기본 원칙

간혹, 메서드 혹은 생성자 파라미터의 조건이 붙을 수가 있다. 예를 들어서, 인덱스 값이 음수이면 안된다던가, 입력으로 들어오는 객체가 널이 되면 안된다는가 하는 코드가 있다. 공개적으로 제공되는 api가 아니라면, 그 값을 이런식으로 검사하는 코드를 넣어도 좋다고 한다.

```java
public class Item49 {
    /**
     * 
     * @param a a는 널이 되면 안된다.
     * @param offset offset은 0보다 크거나 같고, a의 길이보다 작거나 같아야 한다.
     * @param length length는 0보다 크거나 같고, a의 길이 - offset 보다 작거나 같아야 한다.
     */
    private static void sort(long a[], int offset, int length) {
        assert a != null;
        assert offset >= 0 && offset <= a.length;
        assert length >=0 && length <= a.length-offset;

        Arrays.sort(a);
    }
}
```

만약 제약이 있을 경우, 그 제약들을 문서화하고, 코드 시작 부분에 명시적으로 검사하는 것이 좋다.


## 적시에 방어적 복사본을 만들라

자바는 안전한 언어다. 그러나, 개발자는 **클라이언트가 여러분의 코드를 망가뜨리려고 혈안이 되어 있다고 가정하고** 프로그래밍을 해야 한다. 실제 다음의 코드를 살펴보자.

```java
public class Period {
    private final Date start;
    private final Date end;

    public Period(Date start, Date end) {
        this.start = start;
        this.end = end;
    }

    public Date getStart() {
        return new Date();
    }

    public Date getEnd() {
        return end;
    }
}
```

이 경우, `Period`가 생성되면, 그 필드인 `start`, `end`는 바꿀 수 없을 거라 가정한다. 하지만 다음의 경우 불변식이 깨지게 된다.

```java
class Item50Test {
    @Test
    @DisplayName("불변식 테스트")
    public void test01() {
        Period p = new Period(new Date(2020, Calendar.SEPTEMBER, 1), new Date(2020, Calendar.SEPTEMBER, 2));
        Date end = p.getEnd();
        long time = end.getTime();

        end.setTime(time + 50000);
        assertEquals(time, p.getEnd().getTime());
    }
}
```

이렇게 하면 `Date`의 내부 필드가 변하게 되서, 테스트가 실패하게 된다. 어떻게 해결할 수 있을까? `Java 8`에서는 Instant 객체를 쓰면 된다. 여기서는 `Java 8` 이전 버전이라고 가정하고, 방어적인 프로그래밍을 해보자. 

가장 쉬운 방법은 다음과 같다.

```java
public class Period {
    private final Date start;
    private final Date end;

    public Period(Date start, Date end) {
        this.start = start;
        this.end = end;
    }

    public Date getStart() {
        return new Date(start.getTime());
    }

    public Date getEnd() {
        return new Date(end.getTime());
    }
}
```

이런 식으로 하면, 내부 필드는 안전하게 보호된다. 여기서 중요한 것은 `clone` 메서드로 복사해서는 안된다는 것이다. 하지만, 방어적 복사는 비용이 비싸다. 이를 피하려면 확실히 문서화해서, 클라이언트가 할 수 없게끔 만들어주던가 `래퍼 클래스 패턴`을 이용해서 피해를 최소화시키는 방법이 있다.


## 메서드 시그니처를 신중히 설계하라

## 다중정의는 신중히 사용하라

## 가변인수는 신중히 사용하라

## null이 아닌 빈 컬렉션이나 배열을 반환하라

## 옵셔널 반환은 신중히 하라

## 공개된 API 요소에는 항상 문서화 주석을 작성하라


