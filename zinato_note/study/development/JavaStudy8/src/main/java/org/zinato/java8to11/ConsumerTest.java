package org.zinato.java8to11;

import java.util.function.Consumer;

public class ConsumerTest {

  public static void main(String[] args) {
    Consumer<Integer> printT = System.out::println;
    Consumer<Integer> multiplyT = (i) -> System.out.println("info : " + i);
    printT.accept(10);
    printT.andThen(multiplyT).accept(10);
  }

}
