package com.zinato._02_structual_patterns._06_adapter._02_after;

//adaptee
public class Account {

  private String name;

  private String passwword;

  private String email;

  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public String getPasswword() {
    return passwword;
  }

  public void setPasswword(String passwword) {
    this.passwword = passwword;
  }

  public String getEmail() {
    return email;
  }

  public void setEmail(String email) {
    this.email = email;
  }
}
