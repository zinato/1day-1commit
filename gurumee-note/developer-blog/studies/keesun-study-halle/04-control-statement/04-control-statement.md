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

나는 이미 `JUnit4`를 이용하고 있다. 간단히 `JUnit5`로 마이그레이션한다. `build.gradle`을 다음과 같이 수정한다.

```gradle
plugins {
    id 'java'
}

group 'org.example'
version '1.0-SNAPSHOT'

repositories {
    mavenCentral()
}

tasks.withType(JavaCompile) {
    options.compilerArgs += "--enable-preview"
}

tasks.withType(Test) {
    jvmArgs += "--enable-preview"
}

tasks.withType(JavaExec) {
    jvmArgs += '--enable-preview'
}

// 추가된 곳.
test {
    useJUnitPlatform()
}

// 수정된 곳.
dependencies {
    testImplementation 'org.junit.jupiter:junit-jupiter-api:5.3.1'
    testRuntimeOnly 'org.junit.jupiter:junit-jupiter-engine:5.3.1'
}
```

이제 다시 테스트 코드들을 수행하면, 테스트 코드들이 모두 깨진다. 일단 테스트 코드들의 이 부분들을 모두 수정해주어야 한다.

수정 전 코드
```java
// ...
// JUnit4 의존 모듈
import org.junit.Test;

import static org.junit.Assert.assertEquals;
// ...
```

수정 후 코드
```java
// ...
// JUnit5 의존 모듈
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
// ...
```

`assertTrue`, `assertFalse`도 `assertEquals`같이 변경해주면 된다. 그리고 예외의 경우는 다음과 같이 변경 해준다.

수정 전 코드
```java
// ...
    @Test(expected = Exception.class)
    public void exection_test() {
        // 예외가 발생하는 코드
    }
// ...
```

수정 후 코드
```java
// ...
    @Test
    public void exection_test() {
        Assertions.assertThrows(Exception.class, () -> {
            // 예외가 발생하는 코드
        });
    }
// ... 
```

사용 방법은 별로 다르지 않다.

1. 메소드 레벨이 `public`일 것 (`JUnit 5`에서는 `package private` 레벨도 가능하다.)
2. `@Test` 애노테이션이 붙을 것

참고적으로 기존 애노테이션들 중 이름이 변경된 것들이 있는데 다음과 같다.

| JUnit4 | JUnit5 | Description |
| :-- | :-- | :-- |
| @BeforeClass | @BeforeAll | 클래스에 포함된 모든 테스트가 수행되기 전에 실행 |
| @AfterClass | @AfterAll | 클래스에 포함된 모든 테스트가 수행된 후에 실행 |
| @Before | @BeforeEach | 클래스에 포함된 각 테스트가 수행되기 전에 실행 |
| @After | @BeforeEach | 클래스에 포함된 각 테스트가 수행된 후에 실행 |
| @Ignore | @Disable | 테스트 클래스/메소드를 무시 |
| @Category | @Tag | 테스트 필터링 |


또한 `JUnit5`에 추가된 애노테이션은 다음과 같다.

| JUnit5 | Description |
| :-- | :-- |
| @TestFactory | 동적 테스트를 위한 팩토리 메소드 |
| @DisplayName | 테스트 이름을 명시 |
| @Nested | 테스트 클래스 안의 클래스 선언 |
| @ExtendWith | 커스텀 상속 클래스 등록, `RunWith`와 비슷함. |


## 과제 1. live-study 대시 보드 구현

## 과제 2. LinkedList 구현

이전에 공부했던 것이 있어 링크를 남긴다.

* [연결 리스트 구현](https://gurumee92.tistory.com/125)
* [이중 연결 리스트 구현](https://gurumee92.tistory.com/126?category=782305)

오바이긴 하지만, 과제는 "정수형"을 담는 연결 리스트 구현이지만 나는 조금 더 나아가서 "제네릭"을 이용한 단일 연결 리스트를 구현한다. 

먼저 노드이다. 노드는 데이터와, 다음 혹은 이전 노드에 대한 포인터가 필요하다. 단일에 경우는 둘 중 하나만 있으면 되고 보통은 다음 노드를 가리키게끔 구현한다. 현재 내가 가진 지식으로 이 모두를 만족하는 방법은 다음과 같다. 먼저 인터페이스를 선언한다.

```java
public interface ListNode<E extends Comparable<E>> {
    E getData();
    void setData(E data);
    ListNode<E> getNext();
    void setNext(ListNode<E> next);

    ListNode<E> getPrev() throws Exception;
    void setPrev(ListNode<E> prev) throws Exception;
}
```

그리고 단일 노드를 다음과 같이 구현한다.

```java
public class SingleListNode<E extends Comparable<E>> implements ListNode<E> {
    private E data;
    private ListNode<E> next;

    public SingleListNode(E data) {
        this.data = data;
        this.next = null;
    }

    @Override
    public E getData() {
        return data;
    }

    @Override
    public void setData(E data) {
        this.data = data;
    }

    @Override
    public ListNode<E> getNext() {
        return next;
    }

    @Override
    public void setNext(ListNode<E> next) {
        this.next = next;
    }

    @Override
    public ListNode<E> getPrev() throws Exception {
        throw new Exception("can't use method");
    }

    @Override
    public void setPrev(ListNode<E> prev) throws Exception {
        throw new Exception("can't use method");
    }
}
```

이중 연결 리스트의 경우에는 `*Prev` 메소드에서 예외를 발생 안시키게끔 짜면 되지 않을까? 얼추 `ListNode`는 끝났다. 간단한 테스트코드도 짜두자.

```java
class SingleListNodeTest {
    private ListNode<Integer> node;

    @BeforeEach
    public void setUp() {
        Integer data = 5;
        node = new SingleListNode<>(data);
    }

    @Test
    @DisplayName("생성 테스트")
    public void test01() {
        Integer data = 5;
        SingleListNode<Integer> node = new SingleListNode<>(data);
        assertEquals(data, node.getData());
        assertNull(node.getNext());
    }

    @Test
    @DisplayName("set data 테스트")
    public void test02() {
        assertNotNull(node.getData());
        assertNull(node.getNext());

        int newValue = 8;
        node.setData(newValue);
        assertEquals(newValue, node.getData().intValue());
        assertNull(node.getNext());
    }

    @Test
    @DisplayName("setNext 테스트")
    public void test03() {
        assertNotNull(node.getData());
        assertNull(node.getNext());

        ListNode<Integer> newNode = new SingleListNode<>(-1);
        node.setNext(newNode);
        assertNotNull(node.getNext());

        ListNode<Integer> next = node.getNext();
        assertEquals(-1, next.getData().intValue());
        assertNull(next.getNext());
    }

    @Test
    @DisplayName("getPrev exception test")
    public void test04() {
        Assertions.assertThrows(Exception.class, () -> {
           node.getPrev();
        });
    }

    @Test
    @DisplayName("setPrev exception test")
    public void test05() {
        Assertions.assertThrows(Exception.class, () -> {
            ListNode<Integer> newNode = new SingleListNode<>(-1);
            node.setPrev(newNode);
        });
    }
}
```

이제 단일 연결 리스트를 구현한다. `head`, `tail`이 있고 각각 더미 노드에 연결하는 방식을 사용할 것이다. 또한 기선님이 준 요구 사항을 연결리스트, 배열리스트 모두 쓸 수 있도록 다음과 같이 변경하였다.

```java
public interface MyList<E extends Comparable<E>> {
    int size();
    E get(int index);
    void set(int index, E data);
    void add(E data);
    E remove(int index);
    boolean contains(E data);
}
```

인터페이스에 따른 연결 리스트의 코드는 다음과 같다.

```java
public class MySingleLinkedList<E extends Comparable<E>> implements MyList<E> {
    private int size;
    private final ListNode<E> head;
    private final ListNode<E> tail;

    public MySingleLinkedList() {
        size = 0;
        head = new SingleListNode<>(null);
        tail = new SingleListNode<>(null);
        head.setNext(tail);
    }

    @Override
    public int size() {
        return size;
    }

    @Override
    public E get(int index) {
        if (index >= size) {
            throw new IndexOutOfBoundsException();
        }

        ListNode<E> node = head;

        for (int i=0; i<=index; i++) {
            node = node.getNext();
        }

        return node.getData();
    }

    @Override
    public void set(int index, E data) {
        if (index >= size) {
            throw new IndexOutOfBoundsException();
        }

        ListNode<E> node = head;

        for (int i=0; i<=index; i++) {
            node = node.getNext();
        }

        node.setData(data);
    }

    @Override
    public void add(E data) {
        ListNode<E> newNode = new SingleListNode<>(data);
        ListNode<E> node = head;

        while(node.getNext() != tail) {
            node = node.getNext();
        }

        node.setNext(newNode);
        newNode.setNext(tail);
        size += 1;
    }

    @Override
    public E remove(int index) {
        if (index >= size) {
            throw new IndexOutOfBoundsException();
        }

        ListNode<E> prev = head;

        for (int i=0; i<index; i++) {
            prev = prev.getNext();
        }

        ListNode<E> node = prev.getNext();
        E data = node.getData();
        prev.setNext(node.getNext());
        node = null;
        size -= 1;
        return data;
    }

    @Override
    public boolean contains(E data) {
        ListNode<E> node = head.getNext();

        while(node != tail) {
            if (data.equals(node.getData())) {
                return true;
            }

            node = node.getNext();
        }

        return false;
    }
}
```

테스트 코드는 다음과 같이 작성하였다.

```java
class MySingleLinkedListTest {
    private MyList<String> list;

    @BeforeEach
    public void setUp(){
        list = new MySingleLinkedList<>();
    }

    @Test
    @DisplayName("initial size = 0 test")
    public void test01() {
        assertEquals(0, list.size());
    }

    @Test
    @DisplayName("get failed test - empty list")
    public void test02() {
        Assertions.assertThrows(IndexOutOfBoundsException.class, () -> {
            list.get(0);
        });
    }

    private void fixtureAddDataSet() {
        for (int i = 0; i < 5; i++) {
            String input = "res" + i;
            list.add(input);
        }
    }

    @Test
    @DisplayName("add test")
    public void test03() {
        fixtureAddDataSet();

        assertEquals(5, list.size());

        for (int i=0; i<5; i++) {
            String expected = "res" + i;
            assertEquals(expected, list.get(i));
        }
    }

    @Test
    @DisplayName("get failed test - IndexOutOfBoundsException")
    public void test04() {
        fixtureAddDataSet();

        assertEquals(5, list.size());
        Assertions.assertThrows(IndexOutOfBoundsException.class, () -> {
            list.get(6);
        });
    }

    @Test
    @DisplayName("set success test")
    public void test05() {
        fixtureAddDataSet();

        assertEquals(5, list.size());

        for (int i=0; i<5; i++) {
            String changed = "tmp" + i;
            list.set(i, changed);
            assertEquals(changed, list.get(i));
        }
    }

    @Test
    @DisplayName("set failed test - empty list")
    public void test06() {
        Assertions.assertThrows(IndexOutOfBoundsException.class, () -> {
            list.set(0, "here");
        });
    }

    @Test
    @DisplayName("set failed test - IndexOutOfBoundsException")
    public void test07() {
        fixtureAddDataSet();
        Assertions.assertThrows(IndexOutOfBoundsException.class, () -> {
            list.set(6, "here");
        });
    }

    @Test
    @DisplayName("remove test")
    public void test08() {
        fixtureAddDataSet();
        assertEquals(5, list.size());

        for (int i=0; i<5; i++) {
            String remove = list.remove(0);
            assertEquals("res" + i, remove);
        }

        fixtureAddDataSet();
        assertEquals(5, list.size());

        for (int i=0; i<5; i++) {
            String remove = list.remove(list.size()-1);
            assertEquals("res" + (5-(i+1)), remove);
        }

        fixtureAddDataSet();
        assertEquals(5, list.size());

        String remove = list.remove(3);
        assertEquals(4, list.size());
        assertEquals("res3", remove);
        assertEquals("res4", list.get(3));
    }

    @Test
    @DisplayName("remove failed test")
    public void test09() {
        Assertions.assertThrows(IndexOutOfBoundsException.class, () -> {
            list.remove(0);
        });

        fixtureAddDataSet();
        assertEquals(5, list.size());

        Assertions.assertThrows(IndexOutOfBoundsException.class, () -> {
            list.remove(5);
        });
    }

    @Test
    @DisplayName("contains test")
    public void test10() {
        fixtureAddDataSet();
        assertEquals(5, list.size());

        for (int i=0; i<5; i++) {
            String expected = "res"+i;
            assertTrue(list.contains(expected));
        }
    }

    @Test
    @DisplayName("contains failed")
    public void test11() {
        for (int i=0; i<5; i++) {
            String expected = "res"+i;
            assertFalse(list.contains(expected));
        }
    }
}
```

내친 김에 `ArrayList`도 만들자. 배열 기반 Stack/Queue를 만들 때 애를 이용하면 좋을 것이다.

## 과제 3. 배열 기반 Stack 구현

## 과제 4. LinkedList 기반 Stack 구현

## 과제 5. 배열 기반 Queue 구현

## 과제 6. LinkedList 기반 Queue 구현