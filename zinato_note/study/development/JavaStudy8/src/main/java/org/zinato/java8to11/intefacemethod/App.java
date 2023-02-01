package org.zinato.java8to11.intefacemethod;

public class App {

  public static void main(String[] args) {
    Foo foo = new DeafultFoo("zinato");
    foo.printName();
    foo.printNameUpperCase();

    Foo.printFoo();
  }

}
