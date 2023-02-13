package org.zinato.java8to11;

import java.util.function.BinaryOperator;
import java.util.function.UnaryOperator;

public class UnaryOperatorTest {

  public static void main(String[] args) {
    //2개의 인자값이 같을 때 사용할 수 있음
    UnaryOperator<Integer> plus10 = (i) -> i + 10;
    System.out.println("plus10 = " + plus10.apply(10));



  }

}
