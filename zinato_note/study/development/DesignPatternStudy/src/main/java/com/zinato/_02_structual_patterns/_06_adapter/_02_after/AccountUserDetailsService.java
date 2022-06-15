package com.zinato._02_structual_patterns._06_adapter._02_after;

import com.zinato._02_structual_patterns._06_adapter._02_after.security.UserDetails;
import com.zinato._02_structual_patterns._06_adapter._02_after.security.UserDetailsService;

//Adapter
public class AccountUserDetailsService implements UserDetailsService {

  private AccountService accountService;

  public AccountUserDetailsService(AccountService accountService) {
    this.accountService = accountService;
  }

  @Override
  public UserDetails loadUser(String username) {
    Account account = accountService.findAccountByUsername(username);
    return new AccountUserDetails(account);
  }
}
