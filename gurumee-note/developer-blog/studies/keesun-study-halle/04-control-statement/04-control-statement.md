# 스터디 할래 - (4) 제어문

![logo](../logo.png)

> 백기선님의 온라인 스터디 "스터디 할래" 4주차 정리 문서입니다. 

## 목표

자바가 제공하는 제어문을 학습하세요.

**학습할 것 (필수)**

* 선택문
* 반복문

**과제 (선택)**

* JUnit 5 학습
* live-study 대시 보드를 만드는 코드를 작성하세요.
  * [Github Java Library]() 활용
* LinkedList를 구현
* Stack 구현
    * 배열 기반 구현
    * LinkedList 기반 구현
* Queue 구현
    * 배열 기반 구현
    * LinkedList 기반 구현


## 선택문

### if

`if`문은 조건에 따라, 동작 실행을 컨트롤할 수 있다. 가장 간단한 기본 구조는 다음과 같다.

```java
if (조건) {
    구문 1;
} 

구문 2;
```

구문 2는 무조건 실행이 가능하나 조건이 참일 경우 구문 1 -> 구문2 순서대로 실행된다. 코드는 다음과 같이 작성할 수 있다.

```java
public class IfTest {
    @Test
    public void if_true_test() {
        int i = 1;

        if (i == 1) {
            i += 1;
        }

        assertEquals(2, i);
    }

    @Test
    public void if_false_test() {
        int i = 1;

        if (i != 1) {
            i += 1;
        }

        assertEquals(1, i);
    }
}
```

만약 조건에 따라 실행 여부를 결정하고 싶다면 다음 구조를 쓸 수 있다.

```java
if (조건) {
    구문 1;
} else {
    구문 2;
}
```

조건이 참이면 구문1을, 거짓이면 구문2를 실행한다. 코드는 다음과 같이 작성할 수 있다.

```java
public class IfTest {
    // ...

    @Test
    public void if_else_true_test() {
        int i = 1;

        if (i == 1) {
            i += 1;
        } else {
            i -= 1;
        }

        assertEquals(2, i);
    }

    @Test
    public void if_else_false_test() {
        int i = 1;

        if (i != 1) {
            i += 1;
        } else {
            i -= 1;
        }

        assertEquals(0, i);
    }
}
```

또한 여러 조건에 따라 분기할 수 있다. 다음과 같은 구조이다.

```java
if (조건 1) {
    구문 1;
} else if (조건 2) {
    구문 2;
} else if (조건 N) {
    구문 N;
} else {
    구문 N+1;
}
```

코드는 다음과 같이 작성할 수 있다.

```java
public class IfTest {
    // ...
    @Test
    public void multiple_condition_test() {
        int i = 80;
        String res = "";

        if (i > 90) {
            res = "A";
        } else if (i > 80) {
            res = "C";
        } else if (i > 70) {
            res = "D";
        } else {
            res = "E";
        }

        assertEquals("D", res);
    }
}   
```

### switch

`switch`문은 한 변수에 따라 여러 갈래로 나누는 동작을 할 때 쓰는 제어문이다. 코드 구조는 다음과 같다.

```java
switch(variable) {
    case 조건 1: 
        구문 1
        break;
    case 조건 2: 
        구문 2
        break;
    case 조건 N: 
        구문 N
        break;
    default:
        구문 N+1;
        break;
}
```

코드는 다음처럼 쓸 수 있다.

```java
public class SwitchTest {
    @Test
    public void switch_test_a() {
        String res = "A";
        int score = -1;

        switch (res) {
            case "A":
                score = 100;
                break;
            case "B":
                score = 90;
                break;
            case "C":
                score = 80;
                break;
            case "D":
                score = 70;
                break;
            default:
                score = 60;
                break;
        }

        assertEquals(100, score);
    }
}
```

`Java 13`에선 3장에서 봤듯이, `->` 혹은 `yield` 키워드를 통해서 식으로 취급할 수 있다.


## 반복문

### for

`for`는 반복문 중 가장 사랑 받는 반복문이다. 다음과 같은 구조이다.

```java
for (할당문; 조건문; 증감문) {
    구문;
}
```

먼저 변수를 할당하고, 변수에 대한 조건을 검사하고 참이면 구문을 실행 후, 변수를 증가 혹은 감소 시킨다. 이런 식으로 코드를 테스트할 수 있다.

```java
public class ForTest {
    @Test
    public void test_for() {
        int[] arr = {1, 2, 3, 4, 5};
        int res = 0;

        for (int i=0; i<arr.length; i++) {
            res += arr[i];
        }

        assertEquals(15, res);
    }
}
```

### for each

그런데 위의 `for`는 오류를 일으킬 가능성이 있다. 만약 `arr.length`가 아닌 숫자를 썼다고 가정해보자. 5라고 적여야 하는데 실수로 6을 적었다. 어떻게 될까? 

```java
public class ForTest {
    // ...

    @Test(expected = ArrayIndexOutOfBoundsException.class)
    public void test_for_fail() {
        int[] arr = {1, 2, 3, 4, 5};
        int res = 0;

        for (int i=0; i<6; i++) {
            res += arr[i];
        }

        assertEquals(15, res);
    }
}
```

테스트 코드에서 알 수 있듯이 `ArrayIndexOutOfBoundsException`가 발생하게 된다. 이를 근본적으로 해결할 수 있게 향상된 `for`문인 `for each`구문이 있다. 코드 구조는 다음과 같다.

```java
for (변수 : 순회 가능한 객체) {
    구문;
}
```

구조에서 알 수 있듯이 "순회 가능한 객체"라는 제약 사항이 있긴 하다. 코드는 다음과 같이 쓸 수 있다.

```java
public class ForTest {
    // ...
    @Test
    public void test_for_each(){
        int[] arr = {1, 2, 3, 4, 5};
        int res = 0;

        for (int i : arr) {
            res += i;
        }

        assertEquals(15, res);
    }
}
```

### while

`while`문은 조건이 참일 때까지, 구문을 반복하는 것이다. 코드 구조는 다음과 같다.

```java
while(조건) {
    구문;
}
```

코드는 다음과 같이 작성할 수 있다.

```java
public class WhileTest {
    @Test
    public void while_test_01() {
        int i = 1;
        int res = 0;

        while (i < 5) {
            res += i;
            i += 1;
        }

        assertEquals(10, res);
    }
}
```

만약, 처음부터, 조건이 틀리다면, 반복문은 실행되지 않는다. 다음처럼 말이다.

```java
public class WhileTest {
    // ...

    @Test
    public void while_test_02() {
        int i = 5;
        int res = 0;

        while (i < 5) {
            res += i;
            i += 1;
        }

        assertEquals(0, res);
    }
}
```

### do ~ while

`do~while`문 역시 조건이 참일 때까지 코드를 반복한다. 코드 구조는 다음과 같다.

```java
do {
    구문;
} while(조건);
```

코드는 다음과 같이 작성할 수 있다.

```java
public class WhileTest {
    // ...
    @Test
    public void do_while_test_01(){
        int i = 1;
        int res = 0;

        do {
            res += i;
            i += 1;
        } while (i < 5);

        assertEquals(10, res);
    }
}
```

`while`과 달리 `do~while`문은 반드시 한 번은 실행하게 된다. 다음처럼 말이다.

```java
public class WhileTest {
    // ...
    @Test
    public void do_while_test_02(){
        int i = 5;
        int res = 0;

        do {
            res += i;
            i += 1;
        } while (i < 5);

        assertEquals(5, res);
    }
}
```


## 과제 0. JUnit5

## 과제 1. live-study 대시 보드 구현

## 과제 2. LinkedList 구현

## 과제 3. 배열 기반 Stack 구현

## 과제 4. LinkedList 기반 Stack 구현

## 과제 5. 배열 기반 Queue 구현

## 과제 6. LinkedList 기반 Queue 구현