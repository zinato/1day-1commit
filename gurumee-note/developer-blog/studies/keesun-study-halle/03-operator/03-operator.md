# 스터디 할래 - (3) 연산자

![logo](../logo.png)

> 백기선님의 온라인 스터디 "스터디 할래" 3주차 정리 문서입니다.

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
## 관계 연산자
## 논리 연산자
## instanceof
## assignment(=) operator
## 화살표(->) 연산자
## 3항 연산자
## 연산자 우선 순위
## switch 연산자 (Java 13)