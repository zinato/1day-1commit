package org.zinato.java8to11;

public class Consumer {

  public static void main(String[] args) {
    Consumer<Integer> printT = (i) -> System.out.println(i);
  }

}
