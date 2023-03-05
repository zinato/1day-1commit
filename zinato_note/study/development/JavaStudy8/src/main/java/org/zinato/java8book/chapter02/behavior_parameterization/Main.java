package org.zinato.java8book.chapter02.behavior_parameterization;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Main {
  public static void main(String[] args) {
    List<Apple> apples = Arrays.asList(
        new Apple(Color.BLACK, 160),
        new Apple(Color.RED, 130),
        new Apple(Color.RED, 200)
    );

    printPrettyApple(apples, new AppleFancyFormatter());
    printPrettyApple(apples, new AppleSimpleFormatter());
  }

  public static void printPrettyApple(List<Apple> apples, AppleFormatter a) {
    for (Apple apple :  apples) {
      String printSomething = a.accept(apple);
      System.out.println(printSomething);
    }

  }
}
