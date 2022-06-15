package com.zinato._02_structual_patterns._06_adapter._02_after;

public class AccountService {

  public Account findAccountByUsername(String username) {
    Account account = new Account();
    account.setName(username);
    account.setPasswword(username);
    account.setEmail(username);
    return account;
  }

  public void createNewAccount(Account account) {

  }

  public void updateAccount(Account account) {

  }

}
