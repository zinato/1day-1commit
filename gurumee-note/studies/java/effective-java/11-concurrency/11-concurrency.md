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

이렇게 동기화 작업을 하면, 성능이 떨어질 수 있다. 더 빨리 처리하고 싶다면 `volatile` 타입을 쓰는 것도 고려해볼만 하다. 그러나, 가장 좋은 방법은 **가변 데이터는 단일 스레드에서만 쓰는 것**이다. 애초에 동기화같은 복잡한 문제가 사라진다. 스레드 내에서 데이터는 되도록 `불변성(Immutable)`을 지켜주자.

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


> 음 내 컴퓨터에서는 별다른 락이나 에러 없이 0-99까지 계속 수행된다. 왜 그런지 모르겠다. FowardingSet을 안써서 그런가

**응답 불가와 안전 실패를 피하려면, 동기화 메서드나 동기화 블록 안에서 제어를 절대로 클라이언트에 양도하면 안된다.** 예를 들어, 동기화된 영역 안에서 재정의 할 수 있는 메서드를 호출하면 안되며, 클라이언트가 넘겨준 함수 객체를 호출해서도 안된다. 이 예를 살펴보자.

```java
public class ObservableSet<E> extends InstrumentHashSet<E> {
    public ObservableSet(Set<E> s) {
        super(s);
    }

    private final List<SetObserver<E>> observers = new ArrayList<>();

    public void addObserver(SetObserver<E> observer) {
        synchronized (observers) {
            observers.add(observer);
        }
    }

    public boolean removeObserver(SetObserver<E> observer) {
        synchronized (observers) {
            return observers.remove(observer);
        }
    }

    private void notifyElementAdded(E element) {
        synchronized (observers) {
            for (SetObserver<E> observer : observers) {
                observer.added(this, element);
            }
        }
    }

    @Override
    public boolean add (E element) {
        boolean added = super.add(element);

        if (added) {
            notifyElementAdded(element);
        }

        return added;
    }

    @Override
    public boolean addAll(Collection<? extends E> c) {
        boolean result = false;

        for (E element : c) {
            result |= add(element);
        }

        return result;
    }
}
```

여기에 다음 main 코드를 작성해본다.

```java
public static void main(String[] args) {
    ObservableSet<Integer> set = new ObservableSet<>(new HashSet<>());
    set.addObserver((s, e) -> System.out.println(e));

    for (int i=0; i<100; i++) {
        set.add(i);
    }
}
```

0-99까지 잘 출력이 된다. 이제 23일 때, 자기 자신을 제거하는 옵저버를 코드에 넣어본다.


```java
public static void main(String[] args) {
    // ...

    set.addObserver(new SetObserver<Integer>() {
        @Override
        public void added(ObservableSet<Integer> set, Integer element) {
            System.out.println(element);

            if (element == 23) {
                set.removeObserver(this);     
            }
        }
    });
}
```

이렇게 하면, 0-23까지 출력한 후 옵저버가 자신을 구독해지한 후 종료시킨다. 그러나 그렇게 동작하지 않는다. 옵저버의 `added` 호출 시점에, `notifyElementAdded`에서 옵저버들을 순회하고 있기 때문이다. 이 때, `added`는 `ObserverSet`의 `removeObserver`를 호출하고, 이메서드는 다시 `observers.remove`를 호출한다. 여기서 문제가 발생한다. 리스트의 원소를 제거하려는데 리스트를 순회하고 있는 것은 명백한 에러이기 때문이다. 

정리하자면, `notifyElelmentAdded` 메서드에서 수행하는 순회는 동기화 블럭 안에 있기 때문에 동시 수정 발생을 막을 수 있지만, 정작 자신이 콜백을 거쳐 되돌아와 수정하는 것까지는 못막는다.

이번엔 `removeObserver`를 직접 호출하지 말고 `ExecutorService`를 사용해 다른 스레드에게 부탁해보자.

```java
public static void main(String[] args) {
    // ...

    // 아까 같은 위치
    set.addObserver(new SetObserver<Integer>() {
        @Override
        public void added(ObservableSet<Integer> set, Integer element) {
            System.out.println(element);

            if (element == 23) {
                ExecutorService exec = Executors.newSingleThreadExecutor();
                
                try {
                    exec.submit(() -> set.removeObserver(this)).get();
                } catch (Exception e) {
                    e.printStackTrace();
                } finally {
                    exec.shutdown();
                }
            }
        }
    });
}
```

이렇게 하면, 아까의 에러는 사라졌지만, `교착 상태(Dead Lock)`에 빠진다. 백그라운드 스레드가 `set.removeObserver`를 호출하면, 옵저버는 잠금을 시도한다. 그러나 이미 메인 스레드가 락을 쥐고 있는 상태기 때문에, 락을 얻을 수 없다. 동시에, 메인 스레드는 백그라운드 스레드가 옵저버를 제거하기만을 기다리고 있기 때문에 데드락 상황이 벌어진다. 이럴 땐, 동기화 문제가 생키는 코드 블록을 바깥으로 빼주는게 좋다.

`notifyElementAdded`를 다음과 같이 수정한다.

```java
private void notifyElementAdded(E element) {
    List<SetObserver<E>> snapshot = null;
    
    synchronized (observers) {
        snapshot = new ArrayList<>(observers);
    }
    
    for (SetObserver<E> observer : snapshot) {
        observer.added(this, element);
    }
}
```

`snapshot`이 원본 `observers`를 대신해서 해당 작업을 진행한다. 그래서 락이 없이도 안전하게 순회가 된다. 이를 구현한 것이 `CopyOnWriteArrayList`이다. 따라서 코드를 다음과 같이 변경하면, 똑같이 데드락 상황을 피할 수 있다.

```java
public class ObservableSet<E> extends InstrumentHashSet<E> {
    public ObservableSet(Set<E> s) {
        super(s);
    }

    private final List<SetObserver<E>> observers = new CopyOnWriteArrayList<>();

    public void addObserver(SetObserver<E> observer) {
        observers.add(observer);
    }

    public boolean removeObserver(SetObserver<E> observer) {
        return observers.remove(observer);
    }

    private void notifyElementAdded(E element) {
        for (SetObserver<E> observer : observers) {
            observer.added(this, element);
        }
    }

    // ...
}
```

코드가 간단해진다. 이처럼 동시성 문제를 처리할 때는, 자바의 동시성 컬렉션 라이브러리를 잘 사용하는 것이 좋다. 뭐 자기가 구현해도, 메서드를 동기화 영역 바깥에서 호출되게 짜는 것이 좋다. **기본 규칙은 동기화 영역에서 가능한 일을 적게 하는 것이다.** 

동시성 컬렉션 라이브러리는 동시성의 위험으로부터 개발자를 훌륭히 보호해준다. 다만, 성능이 떨어진다. 그래서 동시성 컬렉션 라이브러리와, 컬렉션 라이브러리 사이에서 갈등하고 있다면, 차라리 `문서화`를 통해서 스레드에 안전하지 않다고 기록해두는 것도 하나의 방법이다.


## 스레드보다는 실행자, 태스크, 스트림을 애용하라

일반 스레드를 작성해서 쓰는 것보다, `실행자 프레임워크(Executor Framework)`를 쓰는 것이 더 좋다. 이전 아이템에서 잠깐 쓰는 것을 보았는데, `ExecutorService`는 `Runnable`과 `Callable`을 실행시키고, 우아하게 종료시킨다. 이 녀석의 주요 기능은 다음과 같다.

1. 특정 태스크가 완료되기를 기다린다.
2. 태스크 모음 중 아무것 하나 혹은 모든 태스크가 완료되기를 기다린다.
3. 실행자 서비스가 종료하기를 기다린다.
4. 완료된 태스크들의 결과를 차례로 받는다.
5. 태스크를 특정 시간 혹은 주기적으로 실행하게 한다.

여기서 `Runnable`과 `Callable`를 태스크로 봐도 무방하다. 그냥 여러 스레드를 만들고, 그 스레드들을 관리하기 위한 작업 큐를 만드는 것보다 이 실행자 프레임워크를 사용하는 것이 훨씬 개발자에게 더 좋다. "바퀴를 재발명하지 말자".


## wait와 notify보다는 동시성 유틸리티를 애용하라

자바5 에서는 고수준의 동시성 유틸리티가 도입되었다. 그래서 이전에 쓰이던 `wait-notify` 방식은 매우 어려우니 지양하라고 책은 권고하고 있다. 자바 동시성 유틸리티, `java.util.concurrent`는 크게 3가지로 나눌 수 있다.

1. Executor Framework - 실행자 프레임워크
2. Concurrent Collection - 동시성 컬렉션
3. Synchronizer - 동기화 장치

이전 아이템이서, `Executor Framework`를 살펴봤었다. 이 절에서는 `Concurrent Collection`과 `Synchronizer`에 대해서 살펴 보도록 한다. 

`Concurrent Collection`은 `List`, `Queue`, `Map` 같은 표준 컬렉션 인터페이스에 동시성 처리를 구현하여 만든 고성능 컬렉션이다. 내부에서 동시성 처리를 하기 때문에, 동시성 컬렉션에서 동시성을 무력화하는 것은 불가능하고 외부에서 락을 추가로 사용하면 오히려 속도가 느려진다. 여기서 "동시성 컬렉션이 동시성을 무력화하지 못한다"라는 말은 여러 메서드를 원자적으로 묶어 호출할 수 없다는 말이다.

다음은 `ConcurrentMap`을 사용하는 예제 코드이다.

```java
public class ConcurrentMapExample {
    private static final ConcurrentMap<String, String> map = new ConcurrentHashMap<>();

    public String intern(String s) {
        String result = map.get(s);

        if (result == null) {
            result = map.putIfAbsent(s, s);

            if (result == null) {
                result = s;
            }
        }

        return result;
    }
}
```

실제 테스트해보진 않았지만, `String.intern`보다 빠르다고 한다. 

    참고!
    `String.intern` 메소드 내부에 동시성을 처리하는 로직이 들어서 느리고 메모리 누수가 있다고 한다.

또한, 동시성 컬렉션 이전에 동기화한 컬렉션들이 있다. `Collections.synchronizedMap`이 대표적인 예인데, 이제는 `Concurrent Collection`들을 사용하자. 성능이 훨씬 우수하다.

`Synchronizer`는 스레드가 다른 스레드를 기다릴 수 있게 하여, 서로 작업을 조율한다. 가장 자주 쓰이는 동기화 장치는 `CountDownLatch`, `Semaphore`, 가장 강력한 동기화 장치는 `Phaser`, 그리고 `CyclicBarrier`, `Exchanger` 등이 있다. 책에서는 `CountDownLatch`에 대해서 다룬다.

`CountDownLatch`는 "일회성 장벽"으로써 하나 이상의 스레드가 다른 하나 이상의 스레드 작업이 끝날 때까지 기다리게 하는 것을 강제한다. 다음 예를 살펴보자.

```java
public class CountDownLatchExample {
    public static long time (Executor executor, int concurrency, Runnable action) throws InterruptedException {
        CountDownLatch ready = new CountDownLatch(concurrency); // 동시에 실행할 대수
        CountDownLatch start = new CountDownLatch(1);
        CountDownLatch done = new CountDownLatch(concurrency);

        for (int i=0; i<concurrency; i++) {
            executor.execute(() -> {
                // 스레드가 실행될 때마다 개수 셈.
                ready.countDown();

                try {
                    // wait all worker thread ready
                    start.await();
                    action.run();
                } catch(InterruptedException e) {
                    Thread.currentThread().interrupt();
                } finally {
                    System.out.println("DONE");
                    // 스레드가 끝날 때마다 개수 셈
                    done.countDown();
                }
            });
        }

        ready.await();
        long startNanos = System.nanoTime();
        start.countDown();
        done.await();
        return System.nanoTime() - startNanos;
    }

    public static void main(String[] args) throws InterruptedException {
        ExecutorService service = Executors.newFixedThreadPool(10);
        long t = time(service, 10, () -> {
            System.out.println("HOO~!");
        });
        System.out.println(t);
        service.shutdown();
    }
}
```

굉장히 간단하다. 쓰레드 개수만큼 세는 것이 포인트인 것 같다. 만약, `Executors.newFixedThreadPool`의 매개 변수 nThreads보다 `time`의 매개 변수 concurrency가 더 크다면, 데드락 상황이 벌어져서, 프로그램은 종료하지 않는다. 즉 "nThreads >= concurrency"를 맞춰주어야 한다. (어찌 보면 당연한 얘기)

이 예제는 `CyclicBarrier` 혹은 `Phaser`를 사용하면 훨씬 코드가 명료해진다. 하지만 이해하기는 상대적으로 어렵다. 

이번에는 기존 `wait-notify` 구조에 대해서 간단하게 살펴보자. 다음 규칙에 맞게 짜라.

1. wait 메소드는 대기 반복문(wait loop) 관용구를 사용하라. 반복문 밖에선 사용하지 마라.
2. 대기 조건을 검사하여 조건이 충족하지 않았다면 다시 대기하게 하라.

그러나 다시 대기하게 해도 몇 가지 상황에 의해서 스레드가 깨어날 수도 있다.

* 스레드가 notify를 호출한 다음 대기 중이던 스레드가 깨어나는 사이에 다른 스레드가 락을 얻을 얻을 때 -> 그 락이 보호하는 상태를 변경한다.
* 조건이 만족되지 않았음에도, 다른 스레드가 notify를 호출할 때
* 깨우는 스레드가 지나치게 관대할 때 -> notifyAll을 통해 모두 깨어날 수도 있다.
* 허위각성 시

걍 레거시 아니면 `java.util.concurrent` 쓰자..


## 스레드 안정성 수준을 문서화하라

스레드 안정성 수준을 문서화하는 것은 좋다. (할 수 있나..?) 문서에서 `synchronized` 한정자가 보인다고 하더라도, 스레드에 안전하다고 확신할 수 없다. 메서드 선언에 `synchronized` 한정자를 선언할지는 구현 이슈일 뿐 API에 속하지 않는다. 멀티 스레드 환경에서 API가 안전하게 사용하게 하려면 클래스가 지원하는 스레드 안정성 수준을 정확히 명시해주어야 한다. 다음은 스레드 안정성이 높은 순으로 나열한 것이다.

| 순위 | 내용 | 설명 | 예 |
| :-- | :-- | :-- | :-- |
| 1 | 불변성 객체, Immutable | 상수와 같아서, 외부 동기화가 필요 없다. | Long, String, BigInteger |
| 2 | 무조건적 스레드 안전한 객체, Unconditionally Thread-safe | 내부에서 충실히 동기화하여 별도 외부 동기화 작업 필요x | AtomicLong, ConcurrentHashMap |
| 3 | 조건부 스레드 안전, Conditionally Thread-safe | 일부 메서드는 동시에 사용하려면 외부 동기화가 필요 | Collections.sychronized 래퍼 메서드가 반환하는 컬렉션 |
| 4 | 스레드 안전하지 않음, Not Thead-safe | 각 메서드 호출을 클라이언트가 외부 동기화 메커니즘으로 감싸야 함. | ArrayList, HashMap |
| 5 | 스레드 절대적, Thread-Hostile | 모든 메서드 호출을 동기화로 감싸더라도 멀티 스레드 환경에서 안전하지 않음 | 이건 개발자가 멍청하게 만들 때 |

이런 애노테이션으로 함께 문서화할 수 있다. 

* Immutable
* ThreadSafe
* NotThreadSafe


    참고!
    `lock` 필드는 항상 final로 선언하라. 락 필드의 변경 가능성을 최소화해야 한다.



## 지연 초기화는 신중히 사용하라

지연 초기화는 해당 필드가 필요할 때까지 초기화를 늦추는 기법이다. 주로 최적화 용도에 쓰인다. 그러나 양날의 검이다. 초기화 비용은 줄지만 지연 초기화한 필드에 접근하는 비용은 커지기 때문이다.

그럼에도 지연 초기화가 필요할 때가 있다. 바로 클래스의 인스턴스 중 그 필드를 사용하는 인스턴스의 비율이 낮은 반면, 그 필드를 초기화하는 비용이 클 때이다. 그러나 지연 초기화는 멀티 스레드 환경에서는 정말이지 고통스러운 작업이다. **대부분의 상황에서 일반적인 초기화가 지연 초기화보다 낫다.** 성능 때문에, 어쩔 수 없이 지연 초기화를 해야 한다면, 다음의 지시 사항을 따르자.

* 정적 필드를 지연 초기화할 땐 `지연 초기화 홀더 클래스(Lazy Initialization Holder Class) 관용구`를 사용하자.
    ```java
    private static final FieldHolder {
        static final FieldType field = computeFieldValue();
    }

    private static FieldType getField() { return FieldHolder.field; }
    ```
* 인스턴스 필드를 지연 초기화할 땐 `이중 검사(Double Check) 관용구`를 사용하라
    ```java
    private volatile FieldType field;

    private FieldType getField() {
        FieldType result = field;

        if (result != null) {
            return result;
        }

        synchronized(this) {
            if (result != null) {
                return computeFieldValue();
            }
            return field;
        }
    }
    ```

이중 검사 관용구에는 몇 가지 변형이 있는데, 이 중 단일 검사만 알아보자. 만약 인스턴스 필드가 여러 번 반복해도 된다면, 앞에 null 체크 문을 제거할 수 있다. 또한 `long`, `double`을 제외한 다른 기본 타입이라면, `volatile` 타입을 쓸 필요는 없다.

근데 앵간하면, 지연 초기화는 쓰지 않길 바란다.


## 프로그램의 동작을 스레드 스케줄러에 기대지 말라

프로그램 동작을 스레드 스케줄러에 기대게 되면, 다른 플랫폼에 이식하기 어렵다고 한다. 이식성을 좋게 하려면, **스레드의 평균 수를 프로세스 수보다 지나치게 많도록 하는 것이다.** `java.util.concurrent`를 최대한 쓰고 가능하면, `Thread.yield`는 쓰지 마라. 이식성이 깨지게 되고, 테스트할 방법이 사라진다.