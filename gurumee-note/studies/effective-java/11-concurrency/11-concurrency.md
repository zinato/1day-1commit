# 동시성
===

![대표사진](../intro.png)

> 책 "이펙티브 자바 3판"을 공부하면서 정리한 문서입니다.


## 공유 중인 가변 데이터는 동기화해 사용하라

스레드끼리 공유하는 가변 데이터의 경우 `경쟁 상태(Race Condition)`이라는 문제를 일으켜서 심각한 버그를 만들 수 있다. 이 때 가변 데이터를 `경쟁 상태`를 보호하기 위해서, 즉 동기화를 지원하기 위해 자바는 `synchronized` 키워드를 지원한다. 

책에서 알려주는 동기화의 이점은 크게 다음의 2가지다.

1. 배타적 실행 보장
2. 스레드 사이의 안정적인 통신 보장

> 배타적 실행 = 한 스레드가 변경하는 중이라, 상태가 일관되지 않은 순간 다른 스레드가 보지 못하게 막는 작업.

자바 명세상 `long`, `double` 외 변수를 읽고 쓰는 동작은 `원자적(atomic)`이라고 한다. 그러나 스레드가 필드를 읽을 때, 항상 수정이 완전히 반영된 값을 읽는 것을 보장하지 않는다. 즉 어떤 쓰레드가 읽고 쓰는 것은 원자적으로 동작하나 다른 스레드가 그 수정된 값을 읽는 것은 보장되지 않는다. 쉽게 말해 A 스레드가 공유 변수 x를 1 증가시켰는데, B 스레드에게는 여전히 x가 변하지 않을 수 있다는 것이다.

또한, 책에서는 `Thread.stop`은 사용하지 말라고 경고하고 있다. `Thread.stop`은 안전하지 않아, 동기화가 실패할 확률이 높고 이는, 큰 문제를 발생시킬 수 있다. 그럼 어떻게 스레드를 멈출 수 있을까? 쉽게 생각하기에는 어떤 변수 플래그를 통해서 스레드를 멈추는 방법이 있겠다. 다음 예제 코드처럼 말이다.

```java
public class StopTread {
    private static boolean stopRequested;

    public static void main(String[] args) throws InterruptedException {
        Thread backgroundThread = new Thread(() -> {
            int i =0;
            
            while (!stopRequested) {
                i++;
            }
        });
        backgroundThread.start();

        TimeUnit.SECONDS.sleep(1);
        stopRequested = true;
    }
}
```

그러나 실제 위의 코드는 돌려보면, 무한히 돌아간다. 왜냐하면 동기화가 제대로 이루어지지 않았기 때문이다. 이제 이를 `sychronized` 키워드를 이용해서 동기화를 해보겠다.

```java
public class StopTread {
    private static boolean stopRequested;

    private static synchronized void requestStop() {
        stopRequested = true;
    }

    private static synchronized boolean isStopRequested() {
        return stopRequested;
    }

    public static void main(String[] args) throws InterruptedException {
        Thread backgroundThread = new Thread(() -> {
            int i =0;

            while (!isStopRequested()) {
                i++;
            }
        });
        backgroundThread.start();

        TimeUnit.SECONDS.sleep(1);
        requestStop();
    }
}
```

이런 식으로 가변 데이터인 `stopRequested` 동기화를 걸어준다. 그럼 `requestStop` 메소드와 `isStopRequested` 메소드를 호출할 때마다, `stopRequested`는 스레드로부터 동기화가 보장이 된다. 따라서, 1초 후에 스레드가 정상적으로 종료된다.

이렇게 동기화 작업을 하면, 성능이 떨어질 수 있다. 더 빨리 처리하고 싶다면 `volatile` 타입을 쓰는 것도 고려해볼만 하다. 그러나, 가장 좋은 방법은 **가변 데이터는 단일 스레드에서만 쓰는 것**이다. 애초에 동기화같은 복잡한 문제가 사라진다.

참고로, 신기하게 이렇게 짜면 `synchronized` 없이도 종료가 되더라.

```java
public class StopTread {
    private static boolean stopRequested;

    public static void main(String[] args) throws InterruptedException {
        Thread backgroundThread = new Thread(() -> {
            int i =0;

            while (!stopRequested) {
                System.out.println("Thread " + stopRequested);
                i++;
            }
        });
        backgroundThread.start();

        TimeUnit.SECONDS.sleep(1);
        stopRequested = true;
        System.out.println("MAIN " + stopRequested);
    }
}
```

## 과도한 동기화는 피하라

## 스레드보다는 실행자, 태스크, 스트림을 애용하라

## wait와 notify보다는 동시성 유틸리티를 애용하라

## 스레드 안정성 수준을 문서화하라

## 지연 초기화는 신중히 사용하라

## 프로그램의 동작을 스레드 스케줄러에 기대지 말라