람다와 스트림
===

목차
---
  - [익명 클래스보다는 람다를 사용하라](#익명-클래스보다는-람다를-사용하라)
  - [람다보다는 메서드 참조를 사용하라](#람다보다는-메서드-참조를-사용하라)
  - [표준 함수형 인터페이스를 사용하라](#표준-함수형-인터페이스를-사용하라)
  - [스트림은 주의해서 사용하라](#스트림은-주의해서-사용하라)
  - [스트림에서는 부작용 없는 함수를 사용하라](#스트림에서는-부작용-없는-함수를-사용하라)
  - [반환 타입으로는 스트림보다 컬렉션이 낫다](#반환-타입으로는-스트림보다-컬렉션이-낫다)
  - [스트림 병렬화는 주의해서 적용하라](#스트림-병렬화는-주의해서-적용하라)


## 익명 클래스보다는 람다를 사용하라

이전에는 함수 타입을 표현할 때, 익명 클래스를 사용했다. `Java 8` 이전의 리스트를 정렬하는 코드를 살펴보자.

```java
List<String> words = Arrays.asList("STAT", "START", "HEY");
Collections.sort(words, new Comparator<String>() {
    @Override
    public int compare(String s1, String s2) {
        return Integer.compare(s1.length(), s2.length());
    }
});
```

여기서 `new Comparator<String>() { ... }`가 익명 클래스 객체이다. `Java 8`에서 `Lambda(람다)`라고 해서 이러한 표현식을 간결하게 표현할 수 있게 되었다.

```java
List<String> words = Arrays.asList("STAT", "START", "HEY");
Collections.sort(words, (s1, s2) -> Integer.compare(s1.length(), s2.length()));
```

`람다`는 몇 가지 제약 사항이 있다. **첫 번째 익명 클래스가 구현하는 인터페이스의 메소드가 반드시 1개여야 한다.**

> `Java 8` 부터는 메서드 1개인 인터페이스에 `@FunctionalInterface` 애노테이션을 붙이는게 관례가 되었다.

`Comparator<T>` 인터페이스를 살펴보자.

```java
@FunctionalInterface
public interface Comparator<T> {
  int compare(T o1, T o2);

  // ...
}
```

메서드가 굉장히 많은데, `compare`, `equals`를 제외한 나머지 메서드는 디폴트 메서드이거나 스태틱 메서드이다. `equals`의 경우는 어차피 `Object`의 메서드를 오버라이드하기 때문에, 메서드 1개로 취급할 수 있는 것으로 추정된다.

**두 번째, 코드 자체로 동작히 명확히 설명되지 않거나, 코드 줄 수가 많아지면 람다 사용을 하지 말자.** 현재 `람다`를 사용해서 비교하는 부분은 "문자열의 길이를 비교한다"가 명확하게 표현된다. 하지만, 줄 수가 길어질 수록 명확성은 떨어지게 된다. 책에서는 1줄, 최대 3줄 이하까지 람다로 표현하는 것을 권장하고 있다.

**세 번째, 람다를 직렬화하는 일을 하지 말아라.** 만약 직렬화해야 할 함수 객체가 있다면, `private static` 중첩 클래스의 인스턴스를 사용하는 것이 옳다.


## 람다보다는 메서드 참조를 사용하라

`람다`는 익명 클래스보다 간결하게 표현할 수 있다. 하지만, `람다`보다 `메서드 참조`를 이용하면, 더 간단하게 표현할 수 있다. 이전 절의 `람다`를 활용한 코드를 `메서드 참조` 방식으로 바꿔보겠다.

```java
List<String> words = Arrays.asList("STAT", "START", "HEY");
Collections.sort(words, Comparator.comparingInt(String::length));
```

보다 간결해졌다. 여기서 알아둘 점은 다음과 같다.

* 람다로 할 수 없는 일이면, 메서드 참조로도 할 수 없다.
* 람다보다 명확하지 않거나 짧지 않다면, 메서드 참조를 쓸 필요는 없다.
* 메서드 참조 방식은 다음과 같은 방법이 존재한다.

| 유형| 예 | 같은 기능을 하는 람다 |
| :-- | :-- | :-- |
| static | Integer::parseInt | str -> Integer.parseInt(str) |
| 한정적 instance | Instant.now()::isAfter | Instant then = Instant.now(); t -> thenisAfter(t) |
| 비한정적 instance | String::isLowerCase | str -> str.toLowerCase() |
| class contructor | `Tree<Map<K, V>>::new` | () -> `new TreeMap<K, V>()` |
| array contructor | int[]::new | len -> new int[len] |


참고적으로 `Java 8`의 `List<T>.sort` 디폴트 메서드를 이용하면, 보다 짧은 코드로 표현할 수 있다.

```java
List<String> words = Arrays.asList("STAT", "START", "HEY");
words.sort(Comparator.comparingInt(String::length));
```


## 표준 함수형 인터페이스를 사용하라

`람다`가 지원되면서, "템플릿 메서드 패턴"의 매력이 줄어들었다.

> 템플릿 메서드 패턴이란, 상위 클래스의 기본 메서드를 재정의해 원하는 동작을 구현하는 패턴

이렇게 되면서, API를 작성하는 모범 사례도 바뀌게 되었다. 책에서는 함수형 인터페이스를 작성할 때는 이미 자바 플랫폼의 함수형 인터페이스로 해결이 가능한지 확인하고 쓸 것을 권장하고 있다. `Java 8`의 기본 함수형 인터페이스는 다음과 같다.

| 인터페이스 | 함수 시그니쳐 | 예 |
| :-- | :-- | :-- |
| `UnaryOperator<T>` | `T apply(T t)` | String::toLowerCase |
| `BinaryOperator<T>` | `T apply(T t1, T t2)` | BigInteger::add |
| `Predicate<T>` | `boolean test(T t)` | Collection::isEmpty |
| `Function<T, R>` | `R apply(T t)` | Arrays::asList |
| `Supplier<T>` | `T get()` | Intant::now |
| `Consumer<T>` | `void accept(T t)` | System.out::println |

위의 표에 있는 함수형 인터페이스를 변형해서 총 43개의 함수형 인터페이스가 존재한다. 다음 중 하나 이상을 만족한다면, 전용 함수형 인터페이스를 구현해야 하는지 고민해봐야 한다.

* 자주 쓰이며 이름 자체가 용도를 명확히 설명할 때
* 반드시 따라야 하는 규약이 있을 때
* 유용한 디폴트 메서드를 제공할 때

만들기로 결정했다면 꼭 `@FunctionalInterface`를 붙여주도록 한다.


## 스트림은 주의해서 사용하라

`스트림 API`는 다량의 데이터 처리 작업을 돕고자 `Java 8`에 추가 되었다. API가 제공하는 추상 개념 중 핵심은 다음과 같다.

* 스트림은 데이터 원소의 유한 혹은 무한 시퀀스를 뜻한다.
* 스트림 파이프라인은 스트림에서 시작해, 종단 연산으로 끝나며, 하나 이상의 중간 연산이 있을 수 있다.

이른 바 `Lazy Evaluation`이란 개념이 생긴 것이다. 쉽게 설명하자면, `스트림 API`는 종단 연산이 호출될 때 비로소 그 평가가 이루어진다. 다음의 코드를 보자.

```java
List<Integer> list = Arrays.asList(1, 2, 3, 4, 5); // (1)
int result = list.stream()                         // (2)
        .filter(i -> i % 2 == 1)                   // (3)
        .reduce(0, Integer::sum);                  // (4)
```

코드의 설명은 다음과 같다.

  (1) - { 1, 2, 3, 4, 5 } 정수형, 리스트를 만든다.
  (2) - Collection 을 Stream 으로 변환한다. 1, 2, 3, 4, 5 를 갖는 유한 시퀀스가 만들어진다.
  (3) - 중간 연산이다. 시퀀스에서 홀수인 것들을 걸러낸다. 아직 평가는 이루어지지 않는다.
  (4) - 종단 연산이다. 이 때 평가가 이루어지며, 1, 3, 5의 합을 그 결과로 반환한다.

`스트림 API`는 조심해서 써야 한다. 적절하게 쓰면, 간결하고 가독성이 좋아지는 코드와 함께 성능도 향상 시킬 수 있다. 반면 과도하게 쓰면 읽기 어려울 뿐 아니라 성능도 떨어질 수 있다. 아나그램 예제를 통해서 한 번 더 살펴보자.

> 아나그램이란, 구성하는 알파벳이 같고 순서만 다른 단어를 말한다.

```java
public class AnagramConverter {
    public List<String> convert(int minGroupSize, List<String> words) {
        Map<String, Set<String>> groups = new HashMap<>();
        List<String> result = new ArrayList<>();

        for (String word : words) {
            groups.computeIfAbsent(alphabetize(word), (unused) -> new TreeSet<>()).add(word);
        }

        for (Set<String> group : groups.values()) {
            if (group.size() >= minGroupSize) {
                result.add(group.size() + "-" + group);
            }
        }

        return result;
    }

    private String alphabetize(String s) {
        char[] a = s.toCharArray();
        Arrays.sort(a);
        return new String(a);
    }
}
```

코드를 살펴보면 다음과 같다.

  1. 먼저 최소 그룹 크기인 minGroupSize, 단어 목록인 words를 입력으로 받는다.
  2. words를 순회하여, alphabetize(word)의 결과가 없으면, TreeSet을 생성하고 여기에 단어를 추가한다.
  3. 만들어진 그룹 목록을 순회하여, 그룹 크기가 minGroupSize 보다 그룹의 크기가 크다면, `숫자-[그룹]\n` 문자열로 추가한다. 

만약, `minGroupSize=2, words=["staple", "petals", "result", "tesl"]`의 입력이 들어왔다면, 다음의 결과를 얻게 된다.

1. words 그루핑, 결과 : `{ "aelpst": ["petals", "staple"], "elst": ["tesl"], "elrstu": ["result"] }`
2. minGroup 보다 작으면, 필터링 결과 : `{ "aelpst": ["petals", "staple"] }`
3. 문자열 변환, 결과 : `["2-["petals", "staple"]"]`

이를 이제 `convert` 메서드를 `스트림 API`로 바꿔보자. 

```java
public List<String> convert(int minGroupSize, List<String> words) {
    return words.stream().collect(
            groupingBy(word -> word.chars().sorted()
                    .collect(StringBuilder::new,
                            (sb, c) -> sb.append((char) c),
                            StringBuilder::append)
                    .toString()))
            .values().stream()
            .filter(group -> group.size() >= minGroupSize)
            .map(group -> group.size() + "-" + group)
            .collect(Collectors.toList());
}
```

어떠한가? 굉장히 읽기 복잡하다. 적절하게 기존 코드와 스트림을 적절하게 사용하는 것이 중요하다. 다음 버전을 살펴보자.

```java
public List<String> convert(int minGroupSize, List<String> words) {
    return words.stream()
      .collect(groupingBy(word -> alphabetize(word)))
      .values().stream()
      .filter(group -> group.size() >= minGroupSize)
      .map(group -> group.size() + "-" + group)
      .collect(Collectors.toList());
}
```

`word -> word.chars().sorted() ...` 이 람다 식이 코드를 읽는데 어렵게 만드는 중요 요소다. 어차피 우리는 이 기능을 하는 `alphabetize` 메서드를 만들었다. 람다식보다 `alphabetize` 메서드를 활용하면 보다 간결하게 표현할 수 있다. 책에서는 기존 코드를 스트림으로 리팩토링하되, 새 코드가 나아 보일 때 반영할 것을 권장하고 있다. 더불어서 다음과 같은 상황이라면, 스트림 사용하는데 안성맞춤이라고 한다.

1. 원소들의 시퀀스를 인관되게 변환할 때
2. 원소들의 시퀀스를 필터링할 때
3. 원소들의 시퀀스를 하나의 연산을 사용해 결합할 때
4. 원소들의 시퀀스를 컬렉션에 모을 때
5. 원소들의 시퀀스에서 특정 조건을 만족하는 원소를 찾을 때


## 스트림에서는 부작용 없는 함수를 사용하라

`스트림 API`는 결국 자바에서 함수형 프로그래밍 패러다임을 지원하기 위해서 만들어진 것이라고 볼 수 있다. 함수형 패러다임은 몇 가지 중요 원칙이 있는데 그 중 하나가 부작용이 없는 함수, `Side Effect`가 없는 함수를 사용해야 한다는 것이다. 이를 쉽게 풀어 쓰면, 입력 값이 함수에 들어갔을 때, 변하지 않는 것이다. 쉽게 코드를 살펴보자.

```java
public FunctionExample {
  public static List<String> AddElementList(List<String> list, String elem) {
    list.add(elem);
    return result
  }
}
```

되게 간단하다. `list`에 문자열 `elem`을 집어 넣은 것이다. 그러나 이 코드의 경우 입력 `list`의 불변성이 깨진다. 함수를 거친 후 `list`의 변화가 생긴 것이다. 이를 부작용이 있는 함수라고 말한다. 부작용이 없기 위해서는 이런 메소드를 만들었어야 한다.

```java
public FunctionExample {
    public static List<String> addElement(List<String> list, String elem) {
        List<String> result = new ArrayList<>(list);
        result.add(elem);
        return result;
    }
}
```

위 메소드를 실행시켜도 입력 `list`는 변함이 없다. 이런 것이 불변성을 지키는, 혹은 부작용이 없는 함수라고 말한다. `스트림 API`에서 중간 연산, 종단 연산은 모두 부작용이 없는 함수를 사용해야 불변성을 지킬 수 있다.


## 반환 타입으로는 스트림보다 컬렉션이 낫다

이 부분은 따르자. 어지간하면, 메소드에서 `Stream<T>`를 반환하는 것을 금하자. 책에 핵심 내용만 살펴보자.

1. 스트림으로 처리하기를 원하는 사용자와, 반복자로 처리하길 원하는 사용자 모두 고려하자
2. 반환 원소의 개수가 적다면 표준 컬렉션을 담아 반환하라
3. 스트림과 반복자(Iterable) 어떤 것이 자연스러운지 고민해보라.


## 스트림 병렬화는 주의해서 적용하라

`스트림 API`의 최대 계륵은 바로 병렬화이다. 중간 연산에 `parallel()`을 더하기만 해도 스트림 연산을 병렬로 처리할 수 있다. 하지만, 대부분의 경우에서 성능 향상을 기대하기 어렵다. 최악의 경우 성능은 나빠지고, 결과 자체가 잘못되거나 예상 못한 동작이 생길 수 있다. 정말 잘 쓴다면, 40초 걸릴 것을 4코어 CPU라면 약 10초만에 끝낼 수 있긴 하다. 매력적이긴 하지만 부작용이 심하다. 

만약 병렬화를 쓴다면, 운영 환경에서 충분히 테스트 한 후, 실 서비스 환경으로 옮길 것을 권하고 있다.