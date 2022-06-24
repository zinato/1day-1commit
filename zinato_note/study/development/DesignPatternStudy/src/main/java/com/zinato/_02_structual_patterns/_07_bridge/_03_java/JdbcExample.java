package com.zinato._02_structual_patterns._07_bridge._03_java;

import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;

public class JdbcExample {
  // 브릿지 패턴에서 추상화에 해당함
  public static void main(String[] args) throws ClassNotFoundException {

    // Driver Interface에서 어떤 Driver를 사용할 것인지에 따라 해당 Driver를 사용하고
    // Driver를 변경해도 아래의 DriverManager, Connection, PreparedStatement, ResultSet 모두 그대로 사용 가능
    Class.forName("org.h2.Driver");// 구체적인 구현체

    try (Connection conn = DriverManager.getConnection("jdbc:h2:mem:~/test", "sa", "")) {
      String sql = "Select * from test";
      Statement statement = conn.createStatement();
      statement.execute(sql);

      PreparedStatement statement1 = conn.prepareStatement(sql);
      ResultSet resultSet = statement.executeQuery(sql);
    } catch (SQLException e) {
      throw new RuntimeException(e);
    }
  }

}
