# 3장 템플릿 (2)

![logo](./../logo.png)

    책장속 먼지털기 스터디 5차
    스터디 날짜 : 2020.12.07
    작성 날짜 : 2020.12.06
    페이지 : 240 ~ 277


## 템플릿과 콜백

잠깐 책에 나온 정의를 살펴보자.

템플릿이란?
템플릿은 어떤 목적을 위해 미리 만들어둔 모양이 있는 틀을 말한다. (=`JdbcContext`)

콜백이란?
콜백은 실행되는 것을 목적으로 다른 오브젝트의 메소드에 전달되는 오브젝트를 말한다. (=`StatementStrategy`를 구현한 익명 클래스)

템플릿/콜백 패턴의 일반적인 흐름은 다음과 같다.

![템플릿/콜백 1](06.png)

먼저 클라이언트에서 콜백 객체를 생성한다. `UserDao`의 `deleteAll` 메소드를 보자.

```java
public void deleteAll() throws SQLException {
    StatementStrategy stmt = c -> {
        PreparedStatement ps = c.prepareStatement("delete from users");
        return ps;
    };
    // ...
}
```

위의 코드에서 `deleteAll`은 클라이언트이다. 그리고 콜백인 `StatementStrategy`를 익명 클래스를 생성하는 것을 확인할 수 있다.

![템플릿/콜백 2](07.png)

그 후, 템플릿을 호출하면서 콜백 객체의 참조를 전달한다.

`UserDao`의 `deleteAll`에서 다음 부분이다.

```java
public void deleteAll() throws SQLException {
    StatementStrategy stmt = //..;
    jdbcContext.workWithStatementStrategy(stmt);
}
```

이제 템플릿인 `jdbcContext.workWithStatementStrategy`를 호출한다. 이 때 콜백의 참조 `stmt`를 전달해준다.

![템플릿/콜백 3](08.png)

템플릿은 자신의 코드 동작하면서 필요한 참조 정보들을 만든다. 템플릿인 `JdbcContext` 코드를 살펴보자.

```java
public void workWithStatementStrategy(StatementStrategy stmt) throws SQLException {
    try (Connection c = dataSource.getConnection();
         // ....
    ) {
        // ...
    } catch (SQLException e) {
        throw e;
    }
}
```

이 때 `Connection` 객체를 생성해낸다. 콜백에 필요한 참조를 생성해내고 있다. 

![템플릿/콜백 4](09.png)

그 후, 실행되어야 할 콜백을 호출한다. 계속해서 `JdbcContext`를 보자.

```java
public void workWithStatementStrategy(StatementStrategy stmt) throws SQLException {
    try (// ...
        PreparedStatement ps = stmt.makePreparedStatement(c)
    ) {
        // ...
    } catch (SQLException e) {
        throw e;
    }
}
```

이 떄, `stmt.makePreparedStatement(c)`에서 알 수 있듯이 콜백을 다시 호출한다. 이 때 아까 생성해둔 `Connection` 참조를 전달한다.

![템플릿/콜백 5](10.png)

콜백이 호출되면, 클라이언트 내부의 변수를 직접 참조하면서 작업을 실행한다.

`UserDao.deleteAll`에서는 참조하는 객체가 없다. 바로 다음 단계로 넘어간다. "delete from users" 쿼리를 실행하는 `PreparedStatement`를 생성하고 반환한다.

![템플릿/콜백 6](11.png)

그 결과를 다시 템플릿을 전달한다. 위에서 만든 `PreparedStatement`가 전달된다. 

![템플릿/콜백 7](12.png)

콜백의 결과를 토대로 템플릿은 코드 실행을 계속해서 진행한다. `JdbcContext`에서 다음 부분이다.

```java
public void workWithStatementStrategy(StatementStrategy stmt) throws SQLException {
    try (/* .... */) {
        ps.executeUpdate();
    } catch (SQLException e) {
        throw e;
    }
}
```

콜백의 결과로 전달받은 `PreparedStatement`를 실행하는 것을 볼 수 있다.

![템플릿/콜백 8](13.png)

템플릿 코드가 끝나면 그 결과를 다시 클라이언트에게 전달한다. 이제 템플릿의 결과를 받은 `UserDao.deleteAll`도 마무리가 된다.


## 한 단계 더 나아가서...

`UserDao.deleteAll`은 "delete from users" 쿼리를 전달하고 나머지는 콜백에게 맡긴다. 분면 위 쿼리 말고도 이런 단순 쿼리만을 전달하는 경우가 왕왕 있을 것이다. 이를 위해서 콜백을 다음과 같이 분리해보자. 

```java
@NoArgsConstructor @AllArgsConstructor
@Getter @Setter
public class UserDao {
    // ...
    private void executeSql(final String query) throws SQLException {
        jdbcContext.workWithStatementStrategy(c -> {
            PreparedStatement ps = c.prepareStatement(query);
            return ps;
        });
    }

    public void deleteAll() throws SQLException {
        final String query = "delete from users";
        executeSql(query);
    }

    // ...
}
```

이제 단순 쿼리 전달은 클라이언트 코드에서 쿼리를 생성한 후 `executeSql` 메소드에 전달만 하면 된다. 여기서 더 개선될 점이 있다. 현재 `executeSql`은 `UserDao`만 사용할 수 있다. 아깝지 않은가? 이 콜백을 이용하는 메소드를 조금 더 확장성 있게 사용하기 위해서, 템플릿과 결합하자.

`JdbcContext`을 다음과 같이 수정하자.

```java
@NoArgsConstructor @AllArgsConstructor
@Getter @Setter
public class JdbcContext {
    // ...

    public void executeSql(final String query) throws SQLException {
        workWithStatementStrategy(c -> {
            PreparedStatement ps = c.prepareStatement(query);
            return ps;
        });
    }
}
```

이제 다시 클라이언트인 `UserDao.deleteAll`을 다음과 같이 수정한다.

```java
@NoArgsConstructor @AllArgsConstructor
@Getter @Setter
public class UserDao {
    // ... 

    public void deleteAll() throws SQLException {
        final String query = "delete from users";
        jdbcContext.executeSql(query);
    }

    // ...
}
```

코드를 수정했으니 테스트 코드를 돌려보자. 잘 돌아갈 것이다. 구조적으로 살펴보면 원래는 다음과 같은 구조였다.

![템플릿 콜백 결합 1](./14.png)

클라이언트 내에 콜백이 있었다. 하지만 이제는 콜백이 템플릿과 결합되면서 다음과 같은 구조가 되었다.

![템플릿 콜백 결합 2](./15.png)

보다 응집력이 있는 코드가 만들어졌다. "응집력이 있다"라는 말은 비슷한 일을 하는 코드들이 뭉쳐있다는 뜻이다.


## 콜백의 이해를 위한 예제

## JdbcTemplate 적용

## 스터디원들의 생각 공유

### 나의 질문과 답
   
### 스터디원들의 질문과 답

### 면접 질문으로 생각해볼 것?


