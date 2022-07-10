package com.zinato._02_structual_patterns._06_adapter._02_after;

import com.zinato._02_structual_patterns._06_adapter._02_after.security.UserDetails;

//adapter
public class AccountUserDetails implements UserDetails {

  private Account account;

  public AccountUserDetails(Account account) {
    this.account = account;
  }

  @Override
  public String getUsername() {
    return this.account.getName();
  }

  @Override
  public String getPassword() {
    return this.account.getPasswword();
  }
}
