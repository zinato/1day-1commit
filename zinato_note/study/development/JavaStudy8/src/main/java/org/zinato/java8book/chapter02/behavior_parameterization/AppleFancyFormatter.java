package org.zinato.java8book.chapter02.behavior_parameterization;

public class PrintAppleByFormatter implements AppleFormatter{
  @Override
  public String accept(Apple a) {
    String weight = a.getWeight() > 150 ? " Heavy " : " Lower ";
    String color = a.getColor() + " Apple ";
    return "A" + weight + color;
  }
}
