package com.zinato._01_creational_patterns._01_singleton;

public class RuntimeExample {
  public static void main(String[] args) {
    Runtime runtime = Runtime.getRuntime();
    System.out.println("runtime MaxMemory : " + runtime.maxMemory());
    System.out.println("runtime freeMemory: " + runtime.freeMemory());
  }
}
