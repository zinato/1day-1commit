package org.zinato.java8to11.method_reference;

import java.util.Arrays;
import java.util.function.Function;
import java.util.function.Supplier;
import java.util.function.UnaryOperator;

public class MethodReferenceTest {

  public static void main(String[] args) {
    //static 메소드
    UnaryOperator<String> hi = Greeting::hi;
    //instance method
    Greeting greeting = new Greeting();
    UnaryOperator<String> hello = greeting::hello; // 메소드 참조만 하고 있을 뿐임.
    System.out.println(hello.apply("zinato"));

    //생성자
    Supplier<Greeting> newGreeting = Greeting::new;

    //생성자
    Function<String, Greeting> zinatoGreeting = Greeting::new;
    System.out.println(zinatoGreeting.apply("zinato!!").getName());

    String[] names = {"orange", "car", "apple"};
    Arrays.sort(names, String::compareToIgnoreCase);
    System.out.println(Arrays.toString(names));


  }

}
