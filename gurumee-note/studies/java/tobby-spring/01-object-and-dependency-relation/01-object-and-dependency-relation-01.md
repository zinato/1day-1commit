# 오브젝트와 의존관계 (1)

![logo](./../logo.png)

> 책장속 먼지털기 스터디 1차
> 스터디 날짜 : 2020.11.09
> 작성 날짜 : 2020.11.08 ~ 09
> 페이지 : 53 ~ 102


## 초 간단!하지만, 초 난감한 DAO를 만들어보자

책에 나온대로, id, name, password 세 개의 프로퍼티를 가진 `User`클래스를 만든다.

```java
@NoArgsConstructor @AllArgsConstructor 
@Getter @Setter @ToString @EqualsAndHashCode
public class User {
    private String id;
    private String name;
    private String password;
}
```

> 참고! 자바 빈이란
> 
> 위의 User 클래스는 자바빈이라고 부를 수 있습니다. 간단하게, 다음의 조건이 충족되면 자바빈이라고 말합니다.
> 1) 디폴트 생성자를 갖고 있어야 한다.
> 2) 빈이 노출하는 이름을 가진 속성, 즉 프로퍼티를 get-set으로 제공해주어야 한다.

위 엔티티를 `RDB`의 테이블로 표현하면 다음과 같다.

| Filed Name | Type | Setting |
| :-- | :-- | :-- |
| id | VARCHAR(10) | Primary Key |
| name | VARCHAR(20) | Not Null |
| password | VARCHAR(20) | Not Null |

만약 데이터베이스에 테이블을 만들어야 한다면 다음의 `SQL`을 이용하면 된다.

```sql
create table users (
    id varchar(10) primary key,
    name varchar(20) not null,
    password varchar(20) not null
);
```
이제 이를 데이터베이스에 저장하거나 읽어올 수 있는 DAO를 만들어보자.

> 참고! DAO란?
> 
> Data Access Object의 약자로써, DB를 사용해 데이터를 조회하거나 조작하는 기능을 전담하는 오브젝트를 말합니다. 

코드는 다음과 같다.

```java
import java.sql.*;

public class UserDao {
    public void add(User user) throws ClassNotFoundException, SQLException {
        Class.forName("com.mysql.jdbc.Driver");
        Connection c = DriverManager.getConnection(
                "jdbc:mysql://localhost/springbook", "spring", "book");

        PreparedStatement ps = c.prepareStatement("insert into users(id, name, password) values(?, ?, ?)");
        ps.setString(1, user.getId());
        ps.setString(2, user.getName());
        ps.setString(3, user.getPassword());
        ps.executeUpdate();

        ps.close();
        c.close();
    }

    public User get(String id) throws ClassNotFoundException, SQLException {
        Class.forName("com.mysql.jdbc.Driver");
        Connection c = DriverManager.getConnection(
                "jdbc:mysql://localhost/springbook", "spring", "book");

        PreparedStatement ps = c.prepareStatement("select * from users where id = ?");
        ps.setString(1, id);

        ResultSet rs = ps.executeQuery();
        rs.next();
        User user = new User();
        user.setId(rs.getString("id"));
        user.setName(rs.getString("name"));
        user.setPassword(rs.getString("password"));

        rs.close();
        ps.close();
        c.close();

        return user;
    }
}
```

위 코드에서, JDBC를 이용한 작업의 일반적인 순서를 알 수 있다.

1. 메서드 시그니처에, 예외를 적어준다.(혹은 코드 블럭을 try-catch로 묶어준다)
2. DB 연결을 위한 Connection을 가져온다.
3. SQL을 담은 Statement/PreparedStatement를 만든다.
4. 만들어진 Statement를 실행한다.
5. Get의 경우, ResultSet을 받아서 정보를 저장할 오브젝트로 옮긴다.
6. 작업을 마친 후, 리소스를(ResultSet, Statement, Connection 등) 해제한다.

이제, 작성한 코드를 테스트해보자. 제일 간단하게는 역시 `main`코드를 클래스 내부에 작성하여, 실행시키는 것이다. `UserDao` 클래스에 다음을 작성하자.

```java
public class UserDao {
    // ...

    public static void main(String[] args) throws SQLException, ClassNotFoundException {
        UserDao dao = new UserDao();

        User user = new User("gurumee", "hyunwoo", "ilovespring");
        dao.add(user);
        System.out.println(user.getId() + " register success");

        User user2 = dao.get("gurumee");
        System.out.println(user2.getName() + " " + user2.getPassword());
        System.out.println(user2.getId() + " query success");
    }
}
```

그 후, 코드를 실행해보면, 다음의 결과를 얻을 수 있다.

```bash
gurumee register success
hyunwoo ilovespring
gurumee query success
```

책의 코드는 `MySQL`과, `MySQL`의 `JDBC` 라이브러리인 `mysql:mysql-connector-java`을  사용한다. 로컬/도커를 활용해서 `MySQL`을 설치하고, 아까 살펴봤던 `SQL`을 통해 User 테이블을 만들고 실행을 해야 옳은 결과를 얻을 수 있을 것이다.

그러나 이 코드는 문제점이 많다. 이를 "객체 지향"적으로 리팩토링해보자.


## DAO의 분리

프로그래밍 기초 개념 중 "관심사의 분리"라는 것이 있다. 프로그래밍적으로 말하자면, 관심이 같은 것끼리 모이게 하고, 따로 떨어져 있는 관심끼리는 영향을 주지 않고 분리하는 것을 말한다.

위 `초난감 DAO`에서 세 가지 관심 사항을 발견할 수 있다.

1. DB와 연결을 위한 커넥션을 어떻게 가져오는가
2. DB에 SQL문장을 어떻게 만들고 실행하는가
3. 작업이 끝나면 리소스들을 어떻게 해제할 것인가

이 중 1번에 해당하는 관심사를 분리해보자. 

### 중복 코드를 메서드로 분리하자.

먼저, 연결을 가져오는 중복 코드를 다음과 같이 메서드로 분리한다.

```java
public class UserDao {
    public void add(User user) throws ClassNotFoundException, SQLException {
        Connection c = getConnection();

        // ...
    }

    public User get(String id) throws ClassNotFoundException, SQLException {
        Connection c = getConnection();

        //...
    }

    private Connection getConnection() throws ClassNotFoundException, SQLException {
        Class.forName("com.mysql.jdbc.Driver");
        return DriverManager.getConnection(
                "jdbc:mysql://localhost/springbook", "spring", "book");
    }

    // ...
}
```

위에서 말했듯이 `add`와 `get`에서 중복되는 코드이자, 공통 관심사는 "DB 커넥션을 어떻게 가져오는가" 이다. 그래서 그 공통 코드를 메서드로 분리했다. 

```java
private Connection getConnection() throws ClassNotFoundException, SQLException {
        Class.forName("com.mysql.jdbc.Driver");
        return DriverManager.getConnection(
                "jdbc:mysql://localhost/springbook", "spring", "book");
    }
```

> 참고!
> 
> 인텔리J 기준, cmd + control + m을 누르면 코드를 메서드로 추출할 수 습니다. 아니면, cmd + shift + a를 누른 후 extract method 로 해당 기능을 검색할 수 있습니다.

`UserDao`는 2개의 메소드 밖에 없기 때문에, 티는 안나지만 메서드가 많이 있다고 가정해보자. 이 때, DB 커넥션의 방법을 수정해야 하는 상황이 벌어졌을 때, 메서드를 추출하지 않았다면 어떤 일이 벌어졌겠는가?

메서드를 추출한 지금, 해당 관심사를 담당하는 메서드인 `getConnection`만 수정하면 된다. 하지만 메서드 추출 전이라면, 많은 메서드를 찾아 다니면서 변경을 해야 한다. 이러한 불편함을 제거한 것이다. 

여기에 한 가지 더 리팩토링을 했으면, 반드시 테스트를 통해서 코드가 동작하는지 검증해야 한다. 작성한 main 코드를 돌려보자. 잘 동작하면 다음으로 넘어가자. (돌려보기 전에 DB에 있는 데이터는 삭제하길 바란다.)

### 팩토리 메서드 패턴을 이용해서 UserDao의 getConnection을 분리해보자.

먼저 `UserDao`를 다음과 같이 추상 클래스로 변경한다.

```java
public abstract class UserDao {
    public void add(User user) throws ClassNotFoundException, SQLException {
        // ...
    }

    public User get(String id) throws ClassNotFoundException, SQLException {
        // ...
    }

    public abstract Connection getConnection() throws ClassNotFoundException, SQLException;

    // ...
}
```

`getConnection`을 추상 메서드로 만들어 `UserDao` 클래스를 상속하는 클래스에게 구현을 위임하게 만든다. 책에서처럼 `NUserDao`, `DUserDao`처럼 나눈다고 했을 때, 클래스 구조는 다음과 같이 변한다.

![01](./01.png)

이를 상속하는 `GeneralUserDao`를 만들어보자.

```java
public class GeneralUserDao extends UserDao {
    public Connection getConnection() throws ClassNotFoundException, SQLException {
        Class.forName("com.mysql.jdbc.Driver");
        return DriverManager.getConnection(
                "jdbc:mysql://localhost/springbook", "spring", "book");
    }
}
```

만약, 여러 `UserDao`를 만들어야 한다고 할 때, `GeneralUserDao`처럼, `getConnection` 메서드를 구현하기만 하면 된다. 

이제 다시 테스트다. main 메서드를 다음과 같이 수정한다.

```java
public abstract class UserDao {
    // ...

    public abstract Connection getConnection() throws ClassNotFoundException, SQLException;

     public static void main(String[] args) throws SQLException, ClassNotFoundException {
        UserDao dao = new GeneralUserDao();

        User user = new User("gurumee", "hyunwoo", "ilovespring");
        dao.add(user);
        System.out.println(user.getId() + " register success");

        User user2 = dao.get("gurumee");
        System.out.println(user2.getName() + " " + user2.getPassword());
        System.out.println(user2.getId() + " query success");
    }
}
```

이처럼, `UserDao` 수정 없이 새로운 커넥션을 만드는 서브 클래스들을 만들 수 있다. 조금 더 유연한 코드가 되었다.

이렇게 슈퍼클래스의 기본적인 로직의 흐름을 만들고, 기능의 일부를 추상 메소드나 오버라이딩 가능한 protected 레벨의 메서드를 만든 뒤 서브 클래스의 위임하는 방식을 **템플릿 메서드 패턴**이라고 한다.

또는, 서브 클래스에서 구체적인 오브젝트 생성하게 하는 방식인 **팩토리 메서트 패턴**으로도 볼 수 있다. 팩토리 메서드 패턴의 관점으로 봤을 때, 클래스 다이어그램은 이렇게 구성할 수 있다.

![02](./02.png)

내 코드로는 `GeneralUserDao`는 `Connection G`를 만든다고 볼 수 있다. 위에 나온 디자인 패턴의 설명은 다음을 참고하라.

- [템플릿 메서드 패턴](https://gmlwjd9405.github.io/2018/07/13/template-method-pattern.html)
- [팩토리 메서드 패턴](https://gmlwjd9405.github.io/2018/08/07/factory-method-pattern.html)


## DAO의 확장

## IoC와 스프링 IoC

## 나의 질문과 답


## 스터디원들의 질문과 답




* [스트렛지(전략) 패턴](https://gmlwjd9405.github.io/2018/07/06/strategy-pattern.html)
* [SOLID 원칙](https://johngrib.github.io/wiki/SOLID/)
