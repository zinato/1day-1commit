package org.zinato.java8to11;

import java.util.function.BinaryOperator;

public class BinaryOperatorTest {

  //인자값 2개 리턴값 총 3개가 모두 똑같을 때 사용, 코드가 깔끔해짐
  public static void main(String[] args) {
    BinaryOperator<Integer> plus11 = (i,j) -> i + j + 11;
    System.out.println("plus11 = " + plus11.apply(10,11));
  }

}
