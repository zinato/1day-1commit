package com.zinato._02_structual_patterns._07_bridge._03_java;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class Slf4jExample {

  //로깅 퍼사드를 통해서 로깅을 사용하고 있음 (로깅을 감싸서 만들어 주는 API들을 로깅 퍼사드라고 함)
  //관점에 따라 다른 패턴이라고도 이야기 할 수 있음
  private static Logger logger = LoggerFactory.getLogger(Slf4jExample.class);

  public static void main(String[] args) {
    logger.info("hello logger");
  }
}
