package org.zinato.java8book.chapter02.behavior_parameterization;

public class AppleSimpleFormatter implements AppleFormatter{
  @Override
  public String accept(Apple a) {
    return "An apple weight is " + a.getWeight() +"g";
  }
}
