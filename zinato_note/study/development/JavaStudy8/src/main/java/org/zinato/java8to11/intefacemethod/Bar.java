package org.zinato.java8to11.intefacemethod;

public interface Bar extends Foo{
  //Foo가 제공하는 default 메소드를 제공하고 싶지 않을 때
  @Override
  void printNameUpperCase();
}
