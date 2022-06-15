package com.zinato._02_structual_patterns._06_adapter._03_java;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Enumeration;
import java.util.List;

public class AdapterInJava {

  public static void main(String[] args) {
    //collections
    List<String> strings = Arrays.asList("a", "b", "c"); //adaptee
    Enumeration<String> enumeration = Collections.enumeration(strings); //Adapter
    ArrayList<String> list = Collections.list(enumeration);

    //io
    


  }

}
