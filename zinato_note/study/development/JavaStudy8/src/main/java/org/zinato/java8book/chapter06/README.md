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

### 6.1.2 미리 정의된 컬렉터 P200