package org.zinato.java8book.chapter06;

import java.util.Currency;
import java.util.List;
import java.util.Map;

import static java.util.stream.Collectors.groupingBy;

public class Main {

  Map<Currency, List<Transaction>> transactionByCurrencies =
      transaction.stream().collect(groupingBy(Transaction::getCurrency));
}
