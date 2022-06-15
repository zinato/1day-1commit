package com.zinato._02_structual_patterns._06_adapter._02_after;

import com.zinato._02_structual_patterns._06_adapter._02_after.security.LoginHandler;
import com.zinato._02_structual_patterns._06_adapter._02_after.security.UserDetailsService;

public class App {

  public static void main(String[] args) {
    AccountService accountService = new AccountService();
    UserDetailsService userDetailsService = new AccountUserDetailsService(accountService);
    LoginHandler loginHandler = new LoginHandler(userDetailsService);
    String login = loginHandler.login("zinato", "zinato");
    System.out.println("login = " + login);

  }

}
