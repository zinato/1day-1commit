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

## 반복문

### for

### while

### do ~ while

## 과제 0. JUnit5

## 과제 1. live-study 대시 보드 구현

## 과제 2. LinkedList 구현

## 과제 3. 배열 기반 Stack 구현

## 과제 4. LinkedList 기반 Stack 구현

## 과제 5. 배열 기반 Queue 구현

## 과제 6. LinkedList 기반 Queue 구현