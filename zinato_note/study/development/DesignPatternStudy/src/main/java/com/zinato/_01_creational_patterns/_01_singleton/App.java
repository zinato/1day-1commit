package com.zinato._01_creational_patterns._01_singleton;

import java.io.*;
import java.lang.reflect.Constructor;
import java.lang.reflect.InvocationTargetException;

public class App {
  public static void main(String[] args) throws IOException, ClassNotFoundException {
    Settings settings1 = Settings.getInstance();
    Settings settings2 = null;

    try (ObjectOutput out = new ObjectOutputStream(new FileOutputStream("settings.obj"))){
      out.writeObject(settings1);
    }
    try (ObjectInput in = new ObjectInputStream(new FileInputStream("settings.obj"))) {
      settings2 = (Settings) in.readObject();
    }
    System.out.println(settings1 == settings2);
  }
}
