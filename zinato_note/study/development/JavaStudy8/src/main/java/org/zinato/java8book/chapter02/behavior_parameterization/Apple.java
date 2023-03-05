package org.zinato.java8book.chapter02.behavior_parameterization;

public class Apple {
  Color color;
  int weight;

  public Apple(Color color, int weight) {
    this.color = color;
    this.weight = weight;
  }

  public Color getColor() {
    return color;
  }

  public void setColor(Color color) {
    this.color = color;
  }

  public int getWeight() {
    return weight;
  }

  public void setWeight(int weight) {
    this.weight = weight;
  }
}
