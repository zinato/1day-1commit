package org.zinato.java8to11;

public class Foo {

  public static void main(String[] args) {
    RunSomething runSomething = () -> System.out.println("hello");
    runSomething.doIt();
  }
}
