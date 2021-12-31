package com.zinato._01_creational_patterns._01_singleton;

import java.io.Serializable;

public class Settings implements Serializable {

  private Settings() {}

  private static class SettingsHolder {
    private static final Settings INSTANCE = new Settings();
  }
  public static Settings getInstance() {
    return SettingsHolder.INSTANCE;
  }

  protected Object readResolve() { //역직렬화 할 때 여기를 사용
    return getInstance();
  }
}
