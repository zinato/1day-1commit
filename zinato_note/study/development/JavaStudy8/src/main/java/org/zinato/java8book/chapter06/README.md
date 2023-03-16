# Chapter 06. 스트림으로 데이터 수집 

- 자바 8의 스트림은 데이터 집합을 멋지게 처리하는 게으른 반복자
- collect 다양한 요소 누적방식을 인수로 받아서 스트림을 최종 결과로 도출하는 리듀싱 연산을 수행 가능
- Collection, Collector, collect를 헷갈리지 않도록 주의!

## 6.1 컬렉터란 무엇인가?

### 6.1.1 고급 리듀싱 기능을 수행하는 컬렉터

- 스트림에 collect를 호출하면 스트림의 요소에 리듀싱 연산이 수행된다.
- Collector 인터페이스의 메서드를 어떻게 구현하느냐에 따라 스트림에 어떤 리듀싱 연산을 수행할지 결정된다.
- Collectors 유틀리티 클래스는 자주 사용하는 컬랙터 인스턴스를 손쉽게 생성할 수 있는 정적팩토리 메서드를 제공한다. 
- 가장 많이 사용하는 정적 메서드 -> toList
```java
List<Tranasction> transaction = 
  transactionStream.collect(Collectos.toList());
```

### 6.1.2 미리 정의된 컬렉터 

- Collectors에서 제공하는 메서드의 기능은 크게 세가지로 구분할 수 있다.
  - 스트림 요소를 하나의 값으로 리듀스하고 요약
  - 요소 그룹화
  - 요수 분할

## 6.2 리듀싱과 요약

- 스트림에 있는 객체의 숫자 필드의 합계나 평균 등을 반환하는 연산에도 리듀싱 기능이 자주 사용된다.
- 이러한 연산을 요약 Summarization 연산이라고 부른다.

### 6.2.2 요약 연산

- Collectors 클래스는 Collectors.summingInt라는 객체를 int로 매핑하는 특별한 요약 팩토리 메서드를 제공한다(합계).
- Collectors.summingLong, Collectors.summingDouble 등과 같이 각 형식의 데이터로 합계를 반환한다.
- 합계 : summingInt, summingLong, summingDouble
- 평균값 : averagingInt, averagingLong, averagingDouble
- 최대 : maxBy - Comparator를 인수로 받음
- 최소 : minBy - Comparator를 인수로 받음
- joining 메서드는 내부적으로 StringBuilder를 이용해서 문자열을 하나로 만든다.
  - 스트림의 각 객체에 toString 메서드를 호출해서 추출한 모든 문자열을 하나의 문자열로 연결해서 반환한다.
  ```java
  String shortMenu = menu.stream().map(Dish::getName).collect(joining());
  ```



### 6.2.4 범용 리듀싱 요약 연산

- 모든 컬렉터는 reducing 팩토리 메서드로도 정의할 수 있다.
- 그럼에도 범용 팩토리 메서드 대신 특화된 컬렉터를 사용한 이유는 프로그래밍적 편의성 때문
- 편의성 보다 가독성이 중요할 때도 있다!

```java
int totalCalories = menu.stream().collect(reducing(0, Dish::getCalories, (i,j) -> i+j));
```


### 컬렉션 프레임워크 유연성 : 같은 연산도 다양항 방식으로 수행할 수 있다.

```java 
int totalCalories = me nu.stream().collect(reducing(0, //초기값
Dish::getCalories, //합계 함수
Integer::sum)); //변환함수 
```

- counting 컬렉터도 세 개의 인수를 갖는 reducing 팩토리 메소드를 이용해서 구현 가능하다.

```java
public static <T> Collector<T, ?, Long> counting() {
  return reducing(0L, e -> 1L, Long::sum); 
}
```

- reduce 보다 joining을 사용하는 것이 가독성, 성능 면에서 더 좋음
```
String shortMenu = menu.stream().map(Dish::getName).collect(joining());
String shortMenu = menu.stream().collect(reducing("", Dish::getName, (s1, s2) -> s1+s2));
String shortMenu = menu.stream().collect(reducing((s1, s2) -> s1+s2)).get();
```

## 6.3 그룹화 

- Collectors.groupingBy를 이용해서 쉽게 그룹화 

##### groupingBy와 함께 사용하는 다른 컬렉터 예제

- 메뉴에 있는 모든 요리의 칼로리 합계
```java
Map<Dish.Type, Integer> totalCaloriesByType = 
  menu.stream().collect(groupingBy(Dish::getType), summingInt(Dish::getCalories));
```

## 6.4 분할

- 분할은 분할 함수 (partitioning function)라 불리는 프레디케이트를 분류 함수로 사용하는 특수한 
그룹 기능이다.
- 분할 함수는 boolean을 반환하므로 맵의 키 형식은 Boolean이다.
- 채식 요리와 채식이 아닌 요리 분류
```java
Map<Boolean, List<Dish>> partitionedMenu = 
  menu.stream().collect(partitioningBy(Dish::isVegetarian));

// 참 값의 키로 맵에서 모든 채식 요리를 얻을 수 있다.
        
List<Dish> vegetarianDished = partitionedMenu.get(true);

// 또는 리스트로 생성한 스트림을 프레디케이트로 필터링해서 사용도 가능 
List<Dish> vegetarianDished = 
  menu.stream().filter(Dish::isVegetarian).collect(toList());
```

### 6.4.1 분할의 강점

- 분할 함수가 반환하는 참, 거짓 두 가지 요소의 스트림 리스트를 모두 유지한다는 것이 분할의 장점이다.
- 분할이란 특수한 종류의 그룹화, partitioningBy는 사실 내부적으로 특수한 맵과 두개의 필드로 구현
- partitioningBy 는 반환한 맵 구현은 참과 거짓 두가지 키만 포함하므로 간결하고 효과적

### 6.4.2 숫자를 소수와 비소수로 분할하기
````java
public boolean isPrime(int candidate) {
    return IntStream.range(2, candidate).noneMatch(i -> candidate % i == 0);
  }

  public Map<Boolean, List<Integer>> partitionPrimes(int n) {
    return IntStream.rangeClosed(2, n).boxed()
        .collect(partitioningBy(candidate -> isPrime(candidate)));
  }
````

##### Collectors 클래스의 정적 팩토리 메서드

- counting <Long> : 스트림의 항목 수 계산
```java
long howmanyDishes = menu.stream().collect(counting());
```

- summingInt <Integer> : 스트리밍의 항목에서 정수 프로퍼티 값을 더 함
```java
int totalCalories = menu.stream().collect(summingInt(Dish::getCalories));
```

- summarizingInt <IntSummaryStatistics> : 스트림 내 항목의 최대값, 최솟값, 합계, 평균 등의 정수 정보 통계 수집
```java
IntSummaryStatistics menuStatistics =
        menu.stream().collect(summarizingInt(Dish::getCalories));
```

- joining <String> : 스트림의 각 항목에 toString 메서드를 호출한 결과 문자열 연결
```java
String shorMenu = menu.stream().map(Dish::getName).collect(Collectors.joining(", "));
```

- averagingInt <Double> : 스트림 항목의 정수 프로퍼티의 평균값 계산
- maxBy <Optional<T>> : 주어진 비교자를 이용해서 스트림의 최댓값 요소를 Optional로 감싼 값을 반환,
스트림에 요소가 없을 때는 Optional.empty() 반환
- minBy<Optional<T>> : max 반대 
- reducing : 누적자를 초깃값으로 설정한 다음에 BinaryOperator로 스트림의 각 요소를 반복적으로 누적자와 합쳐 
스트림을 하나의 값으로 리듀싱
```java
int totalCalories = 
  menu.stream().collect(reducing(0, Dish::getCalories, Integer::sum));
```

- collectiongAndThen : 다른 컬렉터를 감싸고 그 결과 변환 함수 적용
```java
int howManyDished = 
  menu.stream().collect(collectingAndThen(toList(), List::size));
```

- groupingBy <Map<K, List<T>> : 하나의 프로퍼티 값을 기준으로 스트림의 항목을 그룹화하며 기준 프로퍼티 값을 결과 맵의 키로 사용
```java
Map<Dish.Type, List<Dish>> dishesByType =
        menu.stream().collect(groupingBy(Dish::getType));
```
- partitioningBy <Map<Boolean, List<T>>> : 프레디케이트를 스트림의 각 항목에 적용한 결과로 항목 분할
```java
Map<Boolean, List<Dish>> vegetarianDishes =
        menu.stream().collect(partitioningBy(Dish::isVegetarian));
```

## 6.5 Collector 인터페이스

```java
public interface Collector<T, A, R> {

  Supplier<A> supplier();

  BiConsumer<A, T> accumulator();

  Function<A, R> finisher();

  BinaryOperator<A> combiner();

  Set<Characteristics> characteristics();
}
```

- T는 수집될 스트림 항목의 제네릭 형식
- A는 누적자, 즉 수집 과정에서 중간 결과를 누적하는 객체의 형식
- R은 수집 연산 결과 객체의 형식 (항상은 아니지만 대개 컬렉션 형식)

### 6.5.1 Collector 인터페이스의 메서드 살펴보기
 
##### supplier 메서드 : 새로운 결과 컨테이너 만들기 

- supplier 메서드는 빈 결과로 이루어진 Supplier를 반환해야한다.
- supplier는 수집 과정에서 빈 누적자 인스턴스를 만드는 파라미터가 없는 함수다. 
- ToListCollector에서 supplier는 아래와 같이 빈 리스트를 반환한다.
```java
public Supplier<List<T>> supplier() {
  return () -> new ArrayList<T>();
}
//생성자 참조 전달 방법도 있음
public Suppler<List<T>> supplier() {
  return ArrayList::new  
}
```

##### accumulator 메서드 : 결과 컨테이너에 요소 추가하기

- accumulator 메서드는 리듀싱 연산을 수행하는 함수를 반환
- 스트림에서 n번 째 요소를 탐색할 때 두 인수, 누적자(n-1개 항목을 수집한 상태)와 n 번째 요소를 함수에 적용
- 함수의 반환값은 void, 요소를 탐색하면서 적용하는 함수에 의해 누적자 내부 상태가 바뀌므로 어떤값일지 단정할 수 없다
- ToListCollector에서 accumulator가 반환하는 함수는 이미 탐색한 항목을 포함하는 리스트에 현재 항목을 추가하는 연산을 수행
```java
public BiConsumer<List<T>, T> accumulator() {
  return (list, item) -> list.add(item);  
}
//메서드 참조
public BiConsumer<List<T>, T> accumulator() {
  reutrn List::add;  
}
```

##### finisher 메서드 : 최종 변환값을결과 컨테이너로 적용하기
- 스트림 탐색을 끝내고 누적자 객체를 최종 결과로 변환하면서 누적 과정을 끝낼 때 호출할 함수를 반환
- 때로는 이미 누적자 객체가 이미 최종 결과인 상황도 있음, 이런 때는 변환과정이 필요하지 않으므로 항등 함수를 반환
```java
public Function<List<T>, List<T>> finisher() {
  return Function.identity();  
}
```

##### combiner 메서드 : 두 결과 컨테이너 병합 P228

