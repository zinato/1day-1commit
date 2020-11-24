# 스터디 할래 - (3) 연산자

![logo](../logo.png)

> 백기선님의 온라인 스터디 "스터디 할래" 3주차 정리 문서입니다. 이 문서는 [자바 오라클 문서](https://docs.oracle.com/javase/tutorial/java/nutsandbolts)를 토대로 만들었습니다.

## 목표

자바가 제공하는 다양한 연산자를 학습하세요.

**학습할 것**

* 산술 연산자
* 비트 연산자
* 관계 연산자
* 논리 연산자
* instanceof
* assignment(=) operator
* 화살표(->) 연산자
* 3항 연산자
* 연산자 우선 순위
* (optional) switch 연산자 (Java 13)


## 산술 연산자

자바의 산술 연산자는 다음과 같다. 

| 연산자 | 설명 |
| :-- | :-- |
| + | 덧셈 연산자 |
| - | 뺄셈 연산자 |
| * | 곱셈 연산자 |
| / | 나눗셈 연산자 |
| % | 나머지 연산자 |

산술 연산자는 "이항 연산자"이다. 기본적으로 다음과 같이 작성한다.

```
# 피연산자1 연산자 피연산자2
a + b
```

여기서 a, b는 피연산자 +는 연산자이다. 이 때 a, b는 실수형, 정수형 타입이다.(코틀린, 스칼라에서는 연산자를 재정의할 수 있기 때문에 더 많은 타입이 들어갈 수 있다.) 각 연산자는 다음과 같이 작성할 수 있다.

**덧셈**
```java
@Test
public void test_add(){
    int a = 5, b = 3;
    // 아래 코드가 덧셈 연산자 쓰는 곳!
    int res = a + b; 
    assertEquals(8, res);
}
```

**뺄셈**
```java
@Test
public void test_sub(){
    int a = 5, b = 3;
    // 아래 코드가 뺼셈 연산자 쓰는 곳!
    int res = a - b;
    assertEquals(2, res);
}
```

**곱셈**
```java
@Test
public void test_mul(){
    int a = 5, b = 3;
    // 아래 코드가 곱셈 연산자 쓰는 곳!
    int res = a * b;
    assertEquals(15, res);
}
```

**나눗셈**
```java
@Test
public void test_div(){
    int a = 5, b = 3;
    // 아래 코드가 나눗셈 연산자 쓰는 곳!
    int res = a / b;
    assertEquals(1, res);
}
```

**나머지**
```java
@Test
public void test_rem(){
    int a = 5, b = 3;
    // 아래 코드가 나머지 연산자 쓰는 곳!
    int res = a % b;
    assertEquals(2, res);
}
```

나눗셈과 나머지 연산은 만약, "피연산자 2"가 0이면 `ArithmeticException`을 발생시킨다.

```java
// 테스트 코드 실행 중 발생하는 `Exception`을 테스트
// 나눗셈과 나머지 연산은 피연산자 2가 0이면 예외가 발생한다.
@Test(expected = ArithmeticException.class)
public void test_div_failed(){
    int a = 5, b = 0;
    // 아래 코드가 나눗셈 연산자 쓰는 곳! 피연산자 2(b)가 0
    int res = a / b;
    assertEquals(1, res);
}

@Test(expected = ArithmeticException.class)
public void test_rem_failed(){
    int a = 5, b = 0;
    // 아래 코드가 나머지 연산자 쓰는 곳! 피연산자 2(b)가 0
    int res = a % b;
    assertEquals(2, res);
}
```


## 비트 연산자

비트 연산자는 정수형 타입에 대해서 "비트 연산"을 한다. 이들은 "산술 연산자"와 마찬가지로 "이항 연산자"이다. 자바에서 제공되는 비트 연산자는 다음과 같다.

| 연산자 | 설명 |
| :-- | :-- |
| & | AND 연산 |
| \| | OR 연산 |
| ^ | XOR 연산 |
| << | left shift 연산 |
| >> | right shift 연산 |

여기서 비트 연산 AND, OR, XOR은 1, 0에 대해서 다음과 같이 연산된다.

**AND**
| - | 0 | 1 |
| :--: | :--: | :--: |
| 0 | 0 | 0 |
| 1 | 0 | 1 |

예를 들어 0001 과 0011 을 AND 연산을 하면, 0001의 결과를 가질 수 있다. 코드로 표현하면 다음과 같다.

```java
@Test
public void test_bit_and(){
    int a = 0x0001;
    int b = 0x0011;
    int res = a & b;
    assertEquals(0x0001, res);
}
```

**OR**
| - | 0 | 1 |
| :--: | :--: | :--: |
| 0 | 0 | 1 |
| 1 | 1 | 1 |

예를 들어 0001 과 0011 을 OR 연산을 하면, 0011의 결과를 가질 수 있다. 코드로 표현하면 다음과 같다.

```java
@Test
public void test_bit_or(){
    int a = 0x0001;
    int b = 0x0011;
    int res = a | b;
    assertEquals(0x0011, res);
}
```

**XOR**
| - | 0 | 1 |
| :--: | :--: | :--: |
| 0 | 0 | 1 |
| 1 | 1 | 0 |

예를 들어 0001 과 0011 을 XOR 연산을 하면, 0010의 결과를 가질 수 있다. 코드로 표현하면 다음과 같다.

```java
@Test
public void test_bit_xor(){
    int a = 0x0001;
    int b = 0x0011;
    int res = a ^ b;
    assertEquals(0x0010, res);
}
```

**Left Shift**
Left Shift 비트 연산자는 비트를 왼쪽으로 1bit씩 민다. 예를 들어 0001이 있다면, 0010이 된다. 코드로 표현하면 다음과 같다.

```java
 @Test
public void test_bit_left_shift(){
    int a = 0x0001;                 // 0x0001 = 00000001
    int res = a << 1;
    assertEquals(0x0002, res);      // 0x0002 = 00000010
}
```

**Right Shift**
Right Shift 비트 연산자는 비트를 오른쪽으로 1bit씩 민다. 예를 들어 0010이 있다면, 0001이 된다. 코드로 표현하면 다음과 같다.

```java
@Test
public void test_bit_right_shift(){
    int a = 0x0002;                 // 0x0002 = 00000010
    int res = a >> 1;
    assertEquals(0x0001, res);      // 0x0001 = 00000001
}
```

Right Shift 비트 연산자는 조금 특별한게 있는데, ">>", ">>>" 이렇게 두 가지가 있다. 각각 ">>"는 sign 비트를 유지하고, ">>>" 0을 넣어준다. 무슨 뜻이냐면 코드를 보면 이해할 수 있다.

```java
@Test
public void test_bit_right_shift2(){
    int a = -2;             // 11111111111111111111111111111110
    int res = a >> 1;       // 11111111111111111111111111111111
    assertEquals(-1, res);  
}
```

오른쪽으로 쉬프트할 때, 가장 최상위의 bit가 유지된다. 반면에 ">>>" 연산자를 써보자.

```java
@Test
public void test_bit_right_shift3(){
    int a = -2;             // 11111111111111111111111111111110
    int res = a >>> 1;      // 01111111111111111111111111111111
    assertEquals(2147483647, res);
}
```

이때는, 상위비트가 유지되지 않고 0으로 채워지며 오른쪽으로 비트가 이동되는 것을 확인할 수 있다.


## 관계 연산자

각 값을 비교하는 것이 이 연산자의 역할이다. 자바가 제공하는 관계 연산자는 다음과 같다. 역시 이항 연산자이다.

| 연산자 | 설명 |
| :-- | :-- |
| == | equal to |
| != | not equal to |
| > | greater than |
| >= | greater than or equal to |
| < | less than |
| <= | less than or equal to |

**== 연산자**
이 연산자는 기본 타입의 경우 값이 같은지 여부를 판단한다. 참조 타입의 경우 같은 참조를 가지고 있는지 여부를 판단한다.

```java
@Test
public void test_equal_to_primitive_type() {
    int a = 5, b = 5;
    assertTrue(a == b);
}

@Test
public void test_equal_to_reference_type() {
    // 같은 값을 가지나, 참조가 다르다.
    Person p1 = new Person(29, "gurumee");
    Person p2 = new Person(29, "gurumee");
    // 이 연산자의 결과는 실패를 가진다.
    assertFalse(p1 == p2);
}
```

**!= 연산자**
이 연산자는 `== 연산자`의 반대이다. 값이 다른지 여부를 판단한다. 역시 기본 타입일 때는, 그 값을 참조 타입일 때는 참조에 대해서 판단한다.

```java
 @Test
public void test_not_equal_to_primitive_type() {
    int a = 5, b = 5;
    // a == b 이기 때문에 실패를 반환한다.
    assertFalse(a != b);
}

@Test
public void test_not_equal_to_reference_type() {
    Person p1 = new Person(29, "gurumee");
    Person p2 = new Person(29, "gurumee");
    // 서로 다른 참조이기 때문에 참을 반환한다.
    assertTrue(p1 != p2);
}
```

**> 연산자**
이 연산자는 피연산자 1이 피연산자 2보다 큰지 여부를 판단한다.

```java
@Test
public void test_greater_than() {
    int a = 7, b = 5;
    assertTrue(a > b);


    a = 5;
    assertFalse(a > b);
}
```

**>= 연산자**
이 연산자는 피연산자 1이 피연산자 2보다 크거나 같은지 여부를 판단한다.
```java
@Test
public void test_greater_than_or_equal_to() {
    int a = 7, b = 5;
    assertTrue(a >= b);

    a = 5;
    assertTrue(a >= b);
}
```

**< 연산자**
이 연산자는 피연산자 1이 피연산자 2보다 작은지 여부를 판단한다.
```java
@Test
public void test_less_than() {
    int a = 3, b = 5;
    assertTrue(a < b);


    a = 5;
    assertFalse(a < b);
}
```

**<= 연산자**
이 연산자는 피연산자 1이 피연산자 2보다 작거나 같은지 여부를 판단한다.
```java
@Test
public void test_less_than_or_equal_to() {
    int a = 3, b = 5;
    assertTrue(a <= b);


    a = 5;
    assertTrue(a <= b);
}
```


## 논리 연산자

논리 연산자는 AND, OR, NOT이며, 참과 거짓에 대해서 판단한다. 자바에서 제공하는 논리 연산자는 다음과 같다.

| 연산자 | 설명 |
| :-- | :-- |
| && | AND |
| \|\| | OR |
| ! | NOT |

**AND 연산자**
논리적으로 피연산자 2개가 모두 참일 때, 참 그 외에는 거짓을 나타낸다. 피연산자1이 거짓일 경우, 연산은 피연산자1만 하고 넘어간다.
| | true  | false |
| :--: | :--: | :--: |
| true | true | false |
| false | false | false |

코드로 보면 다음과 같다.

```java
@Test
public void test_and() {
    int a = 1, b = 2;
    assertTrue(a == 1 && b==2);

    // a == 1이 거짓이 되기 때문에 False를 반환한다.
    a = 3;
    assertFalse(a == 1 && b==2);
}
```

**OR 연산자**
논리적으로 피연산자 2개가 하나라도 참일 때, 참 모두 거짓일 때만 거짓을 나타낸다. 피연산자 1이 참일 경우 연산은 피연산자1만 하고 넘어간다.
| | true | false |
| :--: | :--: | :--: |
| true | true | true |
| false | true | false |

코드로 보면 다음과 같다.

```java
@Test
public void test_or() {
    int a = 1, b = 2;
    assertTrue(a == 1 || b==2);

    // a == 1이 거짓이 되더라도 b == 2를 만족하기 때문에 True를 반환한다.
    a = 3;
    assertTrue(a == 1 || b==2);

    // a == 1, b == 2를 둘다 불만족하기 때문에 False를 반환한다.
    b = 7;
    assertFalse(a == 1 || b==2);
}
```

**NOT 연산자**
피연산자 1개의 논리를 반전시킨다. 참이라면 거짓을, 거짓이라면 참을 반환한다.
| | NOT | 
| :--: | :--: | 
| true | false | 
| false | true |

코드로 보면 다음과 같다.

```java
@Test
public void test_not() {
    int a = 7;
    // a == 7 의 논리를 반전시키기 때문에 False가 나온다.
    // 아래 코드의 경우 a != 7 로 쓰는게 관례이다.
    assertFalse(!(a == 7));
}
```


## instanceof

`instanceof 연산자`는 객체가 어떤 클래스인지 여부를 판단한다. 코드는 다음과 같이 쓸 수 있다.

```java
@Test
public void test_instanceof() {
    String s = "test";
    assertTrue(s instanceof String);
}
```

인터페이스를 구현 혹은 상위 클래스를 상속하는 하위 클래스가 여러 개 일때 각 클래스마다 다른 동작을 부여하고 싶을 때, 좋은 연산자이다. 이런 식으로 말이다.

```java
@Test
public void test_instanceof2(){
    Object [] arr = new Object[] {
            1,
            "test",
            2.0
    };

    for (Object o : arr) {
        if (o instanceof Integer) {
            System.out.println("Integer");
        } else if (o instanceof Double) {
            System.out.println("Double");
        } else if (o instanceof String) {
            System.out.println("String");
        } else {
            System.out.println("I don't know type");
        }
    }
}
```


## assignment(=) operator

이거는 값을 할당하는 연산자이다. 여태까지 코드를 봤을 때 쭉 써왔다.

```
타입 변수 = 값;
```

이런 형태로 쓰는데 코드로 보면 다음과 같다.

```java
int a = 7;
```

위의 경우 a라는 정수형 변수에 7이라는 값을 할당 시킨 것이다.


## 화살표(->) 연산자
## 3항 연산자
## 연산자 우선 순위
## switch 연산자 (Java 13)