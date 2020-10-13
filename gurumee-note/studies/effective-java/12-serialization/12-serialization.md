# 직렬화
===

![대표사진](../intro.png)

> 책 "이펙티브 자바 3판"을 공부하면서 정리한 문서입니다.


## 자바 직렬화의 대안을 찾으라

"직렬화"란 장이란 말이 무색하게 처음부터 직렬화를 쓰지 말 것을 책을 권고하고 있다. 왜냐하면, 직렬화는 매우 위험하기 때문이다. 책에 적힌 `로버트 시커드`의 말을 적어보겠다.

> 자바의 역직렬화는 명백하고 현존하는 위험이다. 이 기술은 지금도 애플리케이션에서 직접 혹은 RMI, JMX, JMS 등의 하부 시스템을 통해 간접적으로 쓰이고 있기 때문이다. 신뢰할 수 없는 스트림을 역직렬화하면 원격 코드 실행, 서비스 거부 등의 공격으로 이어질 수 있다. 잘못한게 아무것도 없는 애플리케이션이라도 이런 공격에 취약해질 수 있다.

이런 위험을 제거하기 위한 제일 좋은 방법은 무엇일까? 바로 "**직렬화를 쓰지 않는 것**"이다.책에서는 다른 대안을 쓸 것을 권고하고 있다. 이를테면, `JSON`, `gRPC(Protobuf)`등 이 대표적인 예이다. 같은 맥락으로 레거시의 경우, **직렬화를 쓰고 있다면, 신뢰할 수 있는 데이터 외에 직렬화/역직렬화하지 말자**.

> 참고! 가젯이란?
>
> 공격자와 보안 전문가들이 자바 라리브러리와 널리 쓰이는 서드파티 라이브러리에서 직렬화 기능 타입들을 연구하여, 역직렬화 과정에서 호출되어 잠재적으로 위험한 동작을 수행하는 메서드를 일컫는다.


## Serializable을 구현할지는 신중히 결정하라

직렬화는 `Serializable` 인터페이스를 `implements`만 선언하기만해도 할 수 있다. 그러나 그 위험성은 앞 절에서 봤듯이 매우 크기 때문에, 구현할 지는 신중히 결정해야 한다. 책에서는 다음의 3가지 이유로 신중히 결정할 것을 권고하고 있다.

1. `Serializable`을 구현하면, 릴리스한 뒤에는 수정하기 어렵다.
2. `Serializable`을 구현하면, 버그와 보안의 구멍이 생길 위험이 높다.
3. `Serializable`을 구현하면, 해당 클래스의 신 버전을 릴리즈할 때 테스트할 것이 늘어난다.

직렬화를 구현하게 되면, 그 클래스는 공개 API가 된다. 널리 퍼진다면, 그 형태로 영원히 지원해야 한다. 또한, 생성자 외에 인스턴스를 만들 수 있기 때문에, 언어의 기본 매커니즘을 우회할 수 있다. 또한 릴리즈된 직렬화를 구현하는 클래스를 버전 별로 직렬화/역직렬화할 수 있는 테스트를 매번 해야 한다. 

책에서는 또한, 상속용으로 설계된 클래스는 `Serializable`을 구현하면 안되며 인터페이스도 `Serializable`을 확장해서도 안된다.


## 커스텀 직렬화 형태를 고려해보라

기본 직렬화 형태를 사용할 수 있는 경우가 몇 가지가 있다. 책에서는 "객체의 물리적 표현과 논리적 내용이 같다면 기본 직렬화 형태를 쓰는 것도 괜찮다"라고 말하고 있다.

```java
public class Name implements Serializable {
    /**
     * 성, null이 아니어야 함.
     * @serial
     */
    private final String lastName;
    /**
     * 이름, null이 아니어야 함.
     * @serial
     */
    private final String firstName;
    /**
     * 중간이름, null이 아니어야 함.
     * @serial
     */
    private final String middleName;

    public Name(String lastName, String firstName, String middleName) {
        this.lastName = lastName;
        this.firstName = firstName;
        this.middleName = middleName;
    }

    // ...
}
```

이름은 논리적으로 이름, 성 중간이름으로 3개의 문자열로 구성된다. 위의 코드는 그것을 그대로 표현하고 있다. 기본 직렬화 형태가 적합하다고 결정했더라도, 불변식 보장과 보안을 위해 `readObject` 메서드로 제공해야 할 때가 있다. 이 때, 해당 필드들은 null이 되지 않는 것을 보장해야 하는데 `Name`은 이를 정확히 따르고 있다. 이제 다른 예를 살표보자.

```java
public class StringList implements Serializable {
    private int size = 0;
    private Entry head = null;

    private static class Entry implements Serializable {
        String data;
        Entry next;
        Entry prev;
    }

    //...
}
```

위의 경우 직렬화 형태에 적합하지 않다. 논리적으로 일련의 문자열을 표현하나, 실제 물리적으로는 이중 연결 리스트로 연결되었다. 즉 물리적 표현과 논리적 표현의 차이가 난다는 것이다. 객체의 물리적 표현과 논리적 표현의 차이가 클 때 기본 직렬화 형태를 사용하면, 크게 4가지 면에서 문제가 생긴다.

1. 공개 API가 현재의 내부 표현 방식에 영원히 묶인다.
2. 너무 많은 공간을 차지할 수 있다.
3. 시간이 오래 걸릴 수 있다.
4. 스택 오버플로를 일으킬 수 있다.

이 경우, 어떻게 고쳐야 할까. 다음은 고친 코드의 예이다.

```java
public class StringList implements Serializable {
    private int size = 0;
    private Entry head = null;
    
    // 이제 직렬화 안됨.
    private static class Entry {
        String data;
        Entry next;
        Entry prev;
    }
    
    // 문자열 추가
    public final void add(String s) {
        //...
    }

    /**
     * 이 {@code StringList} 인스턴스를 직렬화한다.
     * @serialData 이 리스트의 크기를 기록한 후 {@code int}, 이어서 모든 원소를 각각 {@code String} 해서 순서대로 기록한다.
     */
    private void writeObject(ObjectOutputStream s) throws IOException {
        s.defaultWriteObject();
        s.writeInt(size);

        for (Entry e=head; e!=null; e=e.next) {
            s.writeObject(e.data);
        }
    }

    private void readObject(ObjectInputStream s) throws IOException, ClassNotFoundException {
        s.defaultReadObject();
        int numElements = s.readInt();

        for (int i=0; i<numElements; i++) {
            add((String) s.readObject());
        }
    }
}
```

`deafultObject` 메서드를 호출하면, `transient`로 선언하지 않은 모든 인스턴스 필드가 직렬화된다. 해당 객체의 논리적인 상태와 무관할 때만 `transient`를 선언하자. 만약, 직렬화를 결정했다면, 직렬화 가능 클래스 모두에 직렬 버전 UID를 명시적으로 부여한다. 구 버전으로 직렬화된 인스턴스들과의 호환성을 끊으련느 경우를 제외하고는 이 UID를 절대 수정하지 말아야 한다.


## readObject 메서드는 방어적으로 작성하라

`readObject`는 실제적으로 또 다른 public 생성자라고 봐야 한다. 책에서는 **객체를 역직렬화할 때는 클라이언트가 소유해서는 안되는 객체 참조를 갖는 필드를 모두 반드시 방어적으로 복사해야 한다.** 다음은 그 예이다.

```java
private void readObject(ObjectInputStream s) throws IOExecption, ClassNotFoundException {
    s.defaultObject();

    start = new Date(start.getTime());
    end = new Date(end.getTime());

    if (start.compareTo(end) > 0) {
        throw new InvalidObjectException(start + "가 " + end + "보다 늦다.");
    }
}
```

이 때, `clone`을 쓰지 않았음을 주목하자. `readObject` 메서드를 언제 쓰면 좋을까? 판단하기 어렵다면 다음 질문을 던져보라 
    
    `transient` 필드를 제외한 모든 필드의 값을 매개변수로 받아 유효성 검사 없이 필으데 대입하는 public 생성자를 추가해도 괜찮은가?

답이 아니라면, 커스텀 `readObject`를 만들어주어야 한다. 이 때 위처럼, 모든 유효성 검사와 방어적 복사를 수행해야 함은 물론이다. 이 때 좋은 것이 `직렬화 프록시 패턴`이다.


## 인스턴스 수를 통제해야 한다면 readResolve보다는 열거 타입을 사용하라

싱글톤 패턴을 구현한 클래스일지라도 `Serializable`을 구현하게 되면, 더 이상 싱글톤이 아니게 된다. 아래는 이전에 구현했던 싱글톤 패턴의 예제였던 `Elvis` 클래스이다.

```java
public class Elvis implements Serializable {
    public static final Elvis INSTANCE = new Elvis();

    private Elvis() { /* ... */ }
}
```

이렇게 했을 경우, 역직렬화를 통해서 클래스가 만들어질 때 만들어진 인스턴스가 아닌 다른 인스턴스가 만들어진다. 이 때, `readResolve` 메소드를 통해서, 싱글톤 속성을 유지할 수 있다.

```java
public class Elvis implements Serializable {
    public static final Elvis INSTANCE = new Elvis();

    private Elvis() { /* ... */ }

    private Object readResolve() {
        return INSTANCE;
    }

    // ...
}
```

이 때, 모든 필드는 `transient`로 선언해야 한다. 선언하지 않았다면 이렇게 하더라도 클래스는 직렬화/역직렬화로부터 안전하지 않다. 잘 조작된 스트림을 이용하면, 역직렬화하는 시점에 인스턴스의 참조를 훔쳐올 수 있다. 인스턴스 수의 통제가 목적이라면, `readResolve` 메소드보다 열거 타입을 이용하면 더 효과적으로 인스턴스 개수를 통제할 수 있다.

```java
public enum Elvis {
    INSTANCE;

    private String[] favoriteSongs = { "Hound Dog", "Heartbreak Hotel" };
    public void printFavorites() {
        Sytsem.out.println(Arrays.toString(favoriteSongs));
    }
}
```

직렬화 가능 인스턴스 통제 클래스를 작성해야 할 때, 열거 타입이 표현이 불가능하다면, 결국 `readObject`를 고려할 수 밖에 없다. 이 때는 `private` 접근 수준으로 정의하도록 하자.


## 직렬화된 인스턴스 대신 직렬화 프록시 사용을 검토하라

`Serializable`을 구현하는 순간, 생성자 이외에 방법으로 인스턴스를 생성할 수 있게 된다. 즉, 버그와 보안 문제가 일어날 가능성이 높아진다. 이를 그나마 해결해줄 수 있는게, `직렬화 프록시 패턴`이다. 거창한 이름만큼 만드는 게 어렵진 않다. 이전에 구현했던 `Period` 클래스를 예로 살펴보자.

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

먼저 이 클래스 내부에 `Proxy` 클래스를 만들어준다. 직렬화할 필드를 내부 필드로 갖는다.

```java
public class Period {
    // ...
    private static class SerializationProxy implements Serializable {
        private final Date start;
        private final Date end;

        SerializationProxy(Period p) {
            this.start = p.start;
            this.end = p.end;
        }
    }
}
```

이제 이 `Proxy` 클래스에 UID를 만들어준다.

```java
public class Period {
    // ...
    private static class SerializationProxy implements Serializable {
        // ...
        private final static long serialVersionUID = 20202020202L;
    }
}
```

이제 `Period` 클래스의 `writeReplace` 메소드를 구현한다.

```java
public class Period {
    // ...
    private static class SerializationProxy implements Serializable {
        // ...
    }

    private Object writeReplace(){
        return new SerializationProxy(this);
    }
}
```

그리고 `readObject`를 다음과 같이 구현하여 공격자의 공격을 가볍게 막을 수 있다.

```java
public class Period {
    // ...
    private static class SerializationProxy implements Serializable {
        // ...
    }

    // ...

    private void readObject(ObjectInputStream s) throws InvalidObjectException {
        throw new InvalidObjectException("Need Proxy");
    }
}
```

그리고 프록시 클래스 안에 `readResolve`를 구현하면, 역직렬화 시 직렬화 시스템이 직렬화 프록시를 다시 바깥 클래스의 인스턴스로 변환하게 해준다. 

```java
public class Period {
    // ...
    private static class SerializationProxy implements Serializable {
        // ...
        private Object readResolve() {
            return new Period(start, end);
        }
    }

    // ...
}
```

책에서는 제 3자가 확장할 수 없는 클래스를 직렬화해야 한다면, `직렬화 프록시 패턴`을 사용할 것을 권고하고 있다. 다음은 직렬화 프록시 패턴을 사용한 전체 코드이다.

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

    private static class SerializationProxy implements Serializable {
        private final Date start;
        private final Date end;

        SerializationProxy(Period p) {
            this.start = p.start;
            this.end = p.end;
        }

        private final static long serialVersionUID = 20202020202L;

        private Object readResolve() {
            return new Period(start, end);
        }
    }

    private Object writeReplace(){
        return new SerializationProxy(this);
    }

    private void readObject(ObjectInputStream s) throws InvalidObjectException {
        throw new InvalidObjectException("Need Proxy");
    }
}
```