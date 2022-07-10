package com.zinato._02_structual_patterns._06_adapter._02_after.security;

public interface UserDetailsService {

  UserDetails loadUser(String username);

}
