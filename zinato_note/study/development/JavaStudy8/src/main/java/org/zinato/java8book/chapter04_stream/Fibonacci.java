package org.zinato.java8book.chapter04_stream;

import java.util.stream.IntStream;
import java.util.stream.Stream;

public class Fibonacci {

  public static void main(String[] args) {
    Stream.iterate(new int[]{0,1}, t -> new int[]{t[1], t[0] + t[1]})
        .limit(10)
        .forEach(t -> System.out.println("(" + t[0] + "), " + "(" + t[1] + ")"));

    System.out.println("==================");
    Stream.iterate(new int[]{0,1} , t -> new int[]{t[1] , t[0] + t[1]})
        .limit(10)
        .map(t -> t[0] + ", ")
        .forEach(System.out::print);

    IntStream.iterate(0, n -> n <10, n -> n + 4)
        .forEach(System.out::println);

    IntStream.iterate(0, n -> n+4)
        .takeWhile(n -> n < 10)
        .forEach(System.out::println);
  }
}
