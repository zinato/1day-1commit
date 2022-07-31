# Bridge Pattern
- 추상적인 것과 구체적인 것을 분리하여 연결하는 패턴 
- OCP(Open-Closed Principle: 개방 폐쇄 원칙) 디자인 원칙을 지킬 수 있다. 
  - 확장에는 열려있고 변경에는 닫혀있는 
### 구조 

### 개념 

### 장점 
- 추상적인 코드를 구체적인 코드 변경 없이 독립적으로 확장할 수 있다. 
- 추상적인 코드와 구체적인 코드를 분리할 수 있다. 


### 단점
- 계층 구조가 늘어나 복잡도가 증가할 수 있다. 
  - 추상 계층과 Concrete 계층 두 쌍으로 늘어남 

#### 자바와 스프링에서 찾아보는 패턴 

1. JDBC 
   1. 다른 Driver를 쓰더라도 추상화 되어있는 코드를 사용하여 구제적인 코드를 변경없이 사용할 수 있다.
   2. ```
      Connection conn = DriverManager.getConnection("jdbc:h2:mem:~/test", "sa", "")
      Statement statement = conn.createStatement();
      statement.execute(sql);
      ```
2. Sl4j 
   1. LoggerFactory나 Logger는 다른 Logger를 사용하더라도 변경되지 않음
   2. Log4j, LogBack, Log4j2 사용하더라도 추상화 되어있는 것과 구체적인 코드가 분리되어 있기 때문에 그대로 사용 가능
    