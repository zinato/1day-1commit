package org.zinato.java8to11;

@FunctionalInterface
public interface RunSomething {

  /**
   * 함수형 인터페이스는 abstract 메소드가 한개 있어야 한다.
   * 2개가 있으면 컴파일 에러가 발생
   *
   * static, deafult 메소드 사용 가능
   */
  void doIt();

  static void sayHello() {
    System.out.println("say Hello");
  }

  default void saySomething() {
    System.out.println("say something");
  }
}
