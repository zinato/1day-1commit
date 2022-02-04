package com.zinato._01_creational_patterns._02_factory_method._02_after;

import com.zinato._01_creational_patterns._02_factory_method._02_after.Ship;
import com.zinato._01_creational_patterns._02_factory_method._02_after.ShipFactory;

public class Client {

  public static void main(String[] args) {
    Ship whiteship = new WhiteshipFactory().orderShip("Whiteship", "zinato@email.com");
    System.out.println(whiteship);

    Ship blackship = new BlackshipFactory().orderShip("Blackship", "keesun@mail.com");
    System.out.println(blackship);
  }
}
