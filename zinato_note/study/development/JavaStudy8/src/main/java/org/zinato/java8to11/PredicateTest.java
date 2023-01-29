package org.zinato.java8to11;

import java.util.function.Predicate;

public class PredicateTest {

  public static void main(String[] args) {
    Predicate<String> startsWithZ = (s) -> s.startsWith("Z");
    System.out.println(startsWithZ.test("Zinato"));

    //짝수인지
    Predicate<Integer> isEven = (i) -> i % 2 == 0;
    System.out.println(isEven.test(5));
  }

}
