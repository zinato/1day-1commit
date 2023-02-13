package org.zinato.java8to11;

import java.util.function.Supplier;

public class SupplierTest {

  public static void main(String[] args) {
    //인자를 입력할 필요 없이 10을 리턴하는 함수
    Supplier<Integer> get10 = () -> 10;
    System.out.println(get10.get());

  }

}
