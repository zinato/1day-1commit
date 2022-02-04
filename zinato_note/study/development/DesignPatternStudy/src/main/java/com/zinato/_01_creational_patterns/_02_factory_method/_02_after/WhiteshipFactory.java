package com.zinato._01_creational_patterns._02_factory_method._02_after;

public class WhiteshipFactory implements ShipFactory {

  @Override
  public Ship createShip() {
    return new Whiteship();
  }
}
