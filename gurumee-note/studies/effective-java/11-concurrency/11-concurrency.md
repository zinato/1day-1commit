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

## wait와 notify보다는 동시성 유틸리티를 애용하라

## 스레드 안정성 수준을 문서화하라

## 지연 초기화는 신중히 사용하라

## 프로그램의 동작을 스레드 스케줄러에 기대지 말라