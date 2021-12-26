package com.zinato._01_creational_patterns._01_singleton;

import org.springframework.context.ApplicationContext;
import org.springframework.context.annotation.AnnotationConfigApplicationContext;

public class SpringExample {
  public static void main(String[] args) {
    ApplicationContext applicationContext = new AnnotationConfigApplicationContext(SpringConfig.class);
    String hello = applicationContext.getBean("hello", String.class);
    String hello1 = applicationContext.getBean("hello", String.class);

    System.out.println(hello == hello1);
  }
}
