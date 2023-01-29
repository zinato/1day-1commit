package org.zinato.java8to11;

import java.util.function.Consumer;
import java.util.function.IntConsumer;

public class LambdaTest {

  public static void main(String[] args) {
    LambdaTest lambdaTest = new LambdaTest();
    lambdaTest.run();
  }

  private void run() {
    //effective final
    /**
     * Java8부터 지원하는 기능 "사실상" final 변수
     * final 키워드 사용하지 않은 변수를 익명 클래스 구현체 또는 람다에서 참조할 수 있다.
     */
    int baseNumber = 10;

    //로컬 클래스
    class LocalClass {
      void printBaseNumber() {
        int baseNumber = 11;
        System.out.println(baseNumber); //11
      }
    }

    //익명 클래스
    Consumer<Integer> integerConsumer = new Consumer<Integer>() {
      @Override
      public void accept(Integer baseNumber) {
        System.out.println(baseNumber);
      }
    };


    //람다
    /**
     * 익명 클래스 구현체와 달리  shadowing이 일어나지 않음
     * 익명 클래스는 새로 scope을 만들지만, 람다는 감싸고 있는 scope와 같다.
     *
     */
    IntConsumer printInt = (i) -> {
      System.out.println(i + baseNumber);
    };

    printInt.accept(10);
  }

}
