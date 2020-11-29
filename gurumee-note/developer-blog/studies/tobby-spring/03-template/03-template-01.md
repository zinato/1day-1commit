# 3장 템플릿 (1)

![logo](./../logo.png)

    책장속 먼지털기 스터디 4차
    스터디 날짜 : 2020.11.30 (불참)
    작성 날짜 : 2020.11.29 
    페이지 : 209 ~ 240


## 템플릿이란?

책에서 "템플릿"의 정의는 다음과 같다.

바뀌는 성질이 다른 코드 중에서 변경이 겅의 일어나지 않으며 일정한 패턴으로 유지되는 특성을 가진 부분을 자유롭게 변경되는 성질을 가진 부분으로부터 독립시켜서 효과적으로 활용할 수 있는 있도록 하는 방법이다.

쉽게 말하면 변경되지 않는 부분을 남겨두고(템플릿) 변경되는 코드 부분을 독립시키는 것이다.


## UserDao의 문제점?

지난 장들을 거쳐 관심사를 분리하고 의존성 주입 같은 기술들을 적용했음에도 불구하고 `UserDao`는 아직 고쳐야할 점이 남아있다. 바로 "예외 처리"이다.

```java
public class UserDao {
    // ...

    public void deleteAll() throws SQLException {
        Connection c = dataSource.getConnection();

        PreparedStatement ps = c.prepareStatement("delete from users");
        ps.executeUpdate();

        ps.close();
        c.close();
    }

    // ...
}
```

현재 `UserDao`의 `deleteAll` 메소드인데, 무엇인 문제일까? 만약 쿼리를 처리 도중에 오류가 발생하여, `ps.close()`에 도달하기도 전에 종료되었다고 가정하자. 그럼 리소스를 반환하지 않았기 때문에, `PreparedStatement`풀이나, `Connection`풀의 리소스가 남아서 추후에는 리소스가 모자라는 상황이 발생할 수 있다. 

이런 상황을 피하기 위해서 `try-catch-finally` 구문을 사용할 수 있다. 현재 자바에서는 `try-with-resource` 구문을 사용하는 것이 보다 깔끔하게 코드를 처리할 수 있다.

```java
public class UserDao {
    // ...

    public void deleteAll() throws SQLException {
        try(Connection c = dataSource.getConnection();
            PreparedStatement ps = c.prepareStatement("delete from users")) {
            ps.executeUpdate();
        } catch(SQLException e) {
            throw e;
        }
    } 

    // ...
}
```

그러나 "학습"이 목적이므로 `try-catch-finally` 구문으로 한 번 변경해보자.

```java
public class UserDao {
    // ...
    public void deleteAll() throws SQLException {
        Connection c = null;
        PreparedStatement ps = null;
        
        try {
            c = dataSource.getConnection();
            ps = c.prepareStatement("delete from users");
            ps.executeUpdate();
        } catch (SQLException e) {
            throw e;
        } finally {
            if (ps != null) {
                try {
                    ps.close();
                } catch (SQLException e) {
                    
                }
            }
            
            if (c != null) {
                try {
                    c.close();    
                } catch (SQLException e) {
                    
                }
            }
        }
    }
    // ...
}
```

`try-with-resource`구문보다 확실히 코드가 복잡하다. `finally` 블록에서 `ps`와 `c`의 리소스를 반환한다. 이 때 중요한 점은 각각 null 여부인지 체크하고, `close`메소드 역시 `SQLException`이 발생할 수 있으므로 꼭 `try-catch`로 묶어주어야 한다는 점이다. 이제 조회 기능인 `getCount` 메소드에 예외 처리 구문을 추가해보자.

```java
public class UserDao {
    // ...
    public int getCount() throws SQLException {
        Connection c = dataSource.getConnection();

        PreparedStatement ps = c.prepareStatement("select count(*) from users");

        ResultSet rs = ps.executeQuery();
        rs.next();
        int count = rs.getInt(1);

        rs.close();
        ps.close();
        c.close();

        return count;
    }
}
```

`Connection`, `PreparedStatement`외에 `ResultSet` 객체를 추가적으로 더 반환해야 한다. 코드는 다음과 같이 변경된다.

```java
public class UserDao {
    // ...
    public int getCount() throws SQLException {
        Connection c = null;
        PreparedStatement ps = null;
        ResultSet rs = null;

        try {
            c = dataSource.getConnection();
            ps = c.prepareStatement("select count(*) from users");
            rs = ps.executeQuery();
            rs.next();
            int count = rs.getInt(1);
            return count;
        } catch (SQLException e) {
            throw e;
        } finally {
            if (rs != null) {
                try {
                    rs.close();
                } catch (SQLException e) {

                }
            }

            if (ps != null) {
                try {
                    ps.close();
                } catch (SQLException e) {

                }
            }

            if (c != null) {
                try {
                    c.close();
                } catch (SQLException e) {

                }
            }
        }
    }
}
```

더 복잡한 코드가 되었다... 만일 `UserDao`같은 클래스가 수 십개, 수 백개 있다고 해보자. 일일이 `try-catch-finally` 구문으로 다 바꿔줘야 함은 물론 `finally` 블록에서 리소스 반환을 빼먹지 않고 코드를 적었음을 보장해야 한다. 굉장히 고된 작업이 될 것이다. 이를 어떻게 하면 조금 더 깔끔하게, 효율적으로 해결할 수 있을까?


## 템플릿 메소드 패턴? 전략 패턴?



## 전략 패턴 최적화

## 컨텍스트와 DI


## 스터디원들의 생각 공유

### 나의 질문과 답
   
### 스터디원들의 질문과 답



### 면접 질문으로 생각해볼 것?

