package com.zinato._01_creational_patterns._01_singleton;

public enum SettingsWithEnum {
  INSTANCE;

  //생성자 , Getter, Setter 모두 가능 ,
  // reflection에 안전한 코드

  SettingsWithEnum() {
  }


}
