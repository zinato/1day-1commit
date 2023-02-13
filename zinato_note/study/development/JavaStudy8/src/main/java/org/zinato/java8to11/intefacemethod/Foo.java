package org.zinato.java8to11.intefacemethod;

public interface Foo {

  void printName();

  //interface default method

  /**
   * @implSpec 이 구현체는 getName을 대문자로 변환해준다.
   * deafult method는 어떤 기능인지 impleSpec을 통해 적어주는 것이 좋음
   * deafult method 사용하면 implements를 하지 않아도 에러가 나지 않는다.
   * 오버라이드 기능도 가능
   *
   * 예외) Object 메소드는 오버라이드 할 수 없다.
   */
  default void printNameUpperCase() {
    System.out.println(getName().toUpperCase());
  }

  String getName();

  static void printFoo() {
    System.out.println("Static Mehtod Foo!!!");
  }

}
