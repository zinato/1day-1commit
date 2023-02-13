package org.zinato.java8to11.intefacemethod;

public class DeafultFoo implements Foo{

  String name;

  public DeafultFoo(String name) {
    this.name = name;
  }

  @Override
  public void printName() {
    System.out.println("Default Foo");
  }

  @Override
  public String getName() {
    return this.name;
  }
}
