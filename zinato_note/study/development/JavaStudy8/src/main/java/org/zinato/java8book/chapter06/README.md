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

P208