package com.zinato._02_structual_patterns._07_bridge._02_after;

import com.zinato._02_structual_patterns._07_bridge._01_before.Champion;

public class App {

  public static void main(String[] args) {
    Champion kda아리 = new 아리(new KDASkin());
    kda아리.skillQ();
    kda아리.skillR();
  }
}
