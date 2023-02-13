package org.zinato.java8to11;

import java.util.function.Function;

public class Foo {

  public static void main(String[] args) {
    //함수형 인터페이스 default method
    RunSomething runSomething = () -> System.out.println("hello");
    runSomething.doIt();

    Plus10 plus10 = new Plus10();
    System.out.println(plus10.apply(1));

    Function<Integer, Integer> plus11 = (i) -> i+11;
    System.out.println(plus11.apply(1));

    //compose
    Function<Integer, Integer> multiply2 = (i) -> i * 2;
    System.out.println(plus10.compose(multiply2).apply(2)); // 2 * 2 + 10 = 14

    //andThen
    System.out.println(plus10.andThen(multiply2).apply(2)); // 2 + 10 * 2 = 24
  }
}
