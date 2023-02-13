package org.zinato.java8to11.intefacemethod;

import java.util.ArrayList;
import java.util.List;
import java.util.Spliterator;

public class RealEx {

  public static void main(String[] args) {
    List<String> names = new ArrayList<>();
    names.add("a");
    names.add("b");
    names.add("c");
    names.add("d");

    //Iterable Default Method
    System.out.println("ForEach");
    names.forEach(System.out::println);
    System.out.println("------------------");

    //Iterable spliterator
    System.out.println("spliterator");
    Spliterator<String> spliterator = names.spliterator();
    Spliterator<String> spliterator1 = spliterator.trySplit();
    while(spliterator.tryAdvance(System.out::println));
    System.out.println("===============================");
    while (spliterator1.tryAdvance(System.out::println));
    System.out.println("------------------");

    //Collection
    System.out.println("Collection");
    System.out.println("RemoveIf");
    names.removeIf(s -> s.startsWith("a"));
    names.forEach(System.out::println);
    System.out.println("------------------");


    //Comparator
    names.sort(String::compareToIgnoreCase);



  }

}
