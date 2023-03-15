package org.zinato.java8book.chapter06;

import java.util.Arrays;
import java.util.Comparator;
import java.util.Currency;
import java.util.HashMap;
import java.util.IntSummaryStatistics;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

import static java.util.Arrays.asList;
import static java.util.stream.Collectors.collectingAndThen;
import static java.util.stream.Collectors.groupingBy;
import static java.util.stream.Collectors.maxBy;
import static java.util.stream.Collectors.partitioningBy;
import static java.util.stream.Collectors.summarizingInt;
import static java.util.stream.Collectors.toList;
import org.zinato.java8book.chapter06.Dish.Type;

public class Main {

  public static void main(String[] args) {
//    Map<Currency, List<Transaction>> transactionByCurrencies =
//        transaction.stream().collect(groupingBy(Transaction::getCurrency));


//    Map<String, List<String>> dishTags = new HashMap<>();
//    dishTags.put("pork", asList("greasy", "salty"));
//    dishTags.put("beef", asList( "salty", "roasted"));
//    dishTags.put("chicken", asList("fried", "crisp"));
//    dishTags.put("french fries", asList("greasy", "fried"));
//    dishTags.put("rice", asList("light", "natural"));
//    dishTags.put("season fruit", asList("fresh", "natural"));
//    dishTags.put("pizza", asList("tasty", "salty"));
//    dishTags.put("salmon", asList("delicious", "fresh"));


    List<Dish> menu = Arrays.asList(
        new Dish("pork", false, 800, Type.MEAT),
        new Dish("beef", false, 700, Type.MEAT),
        new Dish("chicken", false, 400, Type.MEAT),
        new Dish("french fries", true, 530, Type.OTHER),
        new Dish("rice", true, 350, Type.OTHER),
        new Dish("season fruit", true, 120, Type.OTHER),
        new Dish("pizza", true, 550, Type.OTHER),
        new Dish("prawns", false, 300, Type.FISH),
        new Dish("salmon", false, 450, Type.FISH)
    );


    Map<Boolean, Map<Dish.Type, List<Dish>>> vegetarianDishesByType =
        menu.stream().collect(
            partitioningBy(Dish::isVegetarian,
                groupingBy(Dish::getType))
        );
//    System.out.println(vegetarianDishesByType);


    Map<Boolean, Dish> mostCaloricPartitionedByVegetarian =
        menu.stream().collect(
          partitioningBy(Dish::isVegetarian,
              collectingAndThen(maxBy(Comparator.comparingInt(Dish::getCalories)), Optional::get))
        );

//    System.out.println(mostCaloricPartitionedByVegetarian);

    //summarizingInt = 스트림 내 항목의 최댓값, 최솟값, 합계, 평균등의 정수 정보 통계수집
    IntSummaryStatistics menuStatistics =
        menu.stream().collect(summarizingInt(Dish::getCalories));
//    System.out.println(menuStatistics.getAverage());

    String shorMenu = menu.stream().map(Dish::getName).collect(Collectors.joining(", "));
//    System.out.println(shorMenu);

    int howManyDished =
        menu.stream().collect(collectingAndThen(toList(), List::size));
//    System.out.println(howManyDished);

    Map<Dish.Type, List<Dish>> dishesByType =
        menu.stream().collect(groupingBy(Dish::getType));
//    System.out.println(dishesByType);

    Map<Boolean, List<Dish>> vegetarianDishes =
        menu.stream().collect(partitioningBy(Dish::isVegetarian));

    System.out.println(vegetarianDishes);
    List<Dish> dishes = vegetarianDishes.get(true);
    System.out.println("true : " +  dishes);


  }

  //정수 n을 인수로 받아서 2에서 n까지의 자연수를 소수(Prime)와 비소수(nonPrime)로 나누는 프로그램
  public boolean isPrime(int candidate) {
    return IntStream.range(2, candidate).noneMatch(i -> candidate % i == 0);
  }

  public Map<Boolean, List<Integer>> partitionPrimes(int n) {
    return IntStream.rangeClosed(2, n).boxed()
        .collect(partitioningBy(candidate -> isPrime(candidate)));
  }



}
