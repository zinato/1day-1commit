package com.zinato._02_structual_patterns._07_bridge._01_before;

public class 정복자아리 implements Champion{

  @Override
  public void move() {
    System.out.println("정복자 아리 Move");
  }

  @Override
  public void skillQ() {
    System.out.println("정복자 아리 Q");
  }

  @Override
  public void skillW() {
    System.out.println("정복자 아리 W");
  }

  @Override
  public void skillE() {
    System.out.println("정복자 아리 E");
  }

  @Override
  public void skillR() {
    System.out.println("정복자 아리 R");
  }
}
