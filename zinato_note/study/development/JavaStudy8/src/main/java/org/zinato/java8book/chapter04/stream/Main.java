package org.zinato.java8book.chapter04.stream;

import java.io.IOException;
import java.nio.charset.Charset;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.stream.Stream;

public class Main {

  public static void main(String[] args) {
//    Trader raoul = new Trader("Raoul", "Cambridge");
//    Trader mario = new Trader("Mario", "Milan");
//    Trader alan = new Trader("Alan", "Cambridge");
//    Trader brian = new Trader("Brian", "Cambridge");
//
//    List<Transaction> transactions = Arrays.asList(
//        new Transaction(brian, 2011, 300),
//        new Transaction(raoul, 2012, 1000),
//        new Transaction(raoul, 2011, 400),
//        new Transaction(mario, 2012, 710),
//        new Transaction(mario, 2012, 700),
//        new Transaction(alan, 2012, 950)
//    );
//    //1. 2011년에 일어난 모든 트랜잭션을 찾아 값을 오름차순으로 정리하시오.
//    List<Transaction> collects = transactions.stream()
//        .filter(t -> t.getYear() == 2011)
//        .sorted(Comparator.comparing(Transaction::getValue))
//        .collect(Collectors.toList());
//    System.out.println(collects);
//    //2. 거래자가 근무하는 모든 도시를 중복 없이 나열하시오.
//    List<String> cities = transactions.stream()
//        .map(t -> t.getTrader().getCity())
//        .distinct()
//        .collect(Collectors.toList());
//    System.out.println(cities);
//    //3. 케임브리지에서 근무하는 모든 거래자를 찾아서 이름순으로 정렬하시오.
//    List<Trader> tradersInCambridge = transactions.stream()
//        .map(Transaction::getTrader)
//        .filter(t -> t.getCity().equals("Cambridge"))
//        .sorted(Comparator.comparing(Trader::getName))
//        .distinct()
//        .collect(Collectors.toList());
//    System.out.println(tradersInCambridge);
//    //4. 모든 거래자의 이름을 알파벳순으로 정렬해서 반환하시오.
//    String namesByAlphabet = transactions.stream()
//        .map(t -> {
//          System.out.println("trander string : " + t.getTrader().getName());
//          return t.getTrader().getName();
//        })
//        .distinct()
//        .sorted()
//        .reduce("", (a, b) -> a + b);
//    System.out.println(namesByAlphabet);
//    //5. 밀라노에 거래자가 있는가?
//    boolean isMilanAnyMatch = transactions.stream()
//        .anyMatch(t -> t.getTrader().getCity().equals("Milan"));
//    System.out.println(isMilanAnyMatch);
//    //6. 케임브리지에 거주하는 거래자의 모든 트랜잭션 값을 출력하시오.
//    transactions.stream()
//        .filter(t -> t.getTrader().getCity().equals("Cambridge"))
//        .map(Transaction::getValue)
//        .forEach(System.out::println);
//    //7. 전체 트랜잭션 중 최대값은 얼마인가?
//    Integer maxValue = transactions.stream()
//        .map(Transaction::getValue)
//        .reduce(0, Integer::max);
//    System.out.println("maxValue : " + maxValue);
//    //8. 전체 트랜잭션 중 최솟값은 얼마인가?
//    Optional<Integer> minValue = transactions.stream()
//        .map(Transaction::getValue)
//        .reduce(Integer::min);
//    System.out.println("minValue : " + minValue);


    Stream<String> homeValueStream = Stream.ofNullable(System.getProperty("home"));
    homeValueStream.forEach(System.out::println);

    long uniqueWords = 0;
    try(Stream<String> lines = Files.lines(Paths.get("data.txt"), Charset.defaultCharset())) {
      uniqueWords = lines.flatMap(line -> Arrays.stream(line.split(" ")))
          .distinct()
          .count();
//      System.out.println("uniqueWords : " + uniqueWords);
    } catch (IOException e) {
      e.printStackTrace();
    }


    Stream.iterate(0, n -> n+2)
        .limit(10)
        .forEach(System.out::println);

  }


}


