# 2장 테스트

![logo](./../logo.png)

    책장속 먼지털기 스터디 3차
    스터디 날짜 : 2020.11.23 (불참)
    작성 날짜 : 2020.11.26 - 2020.11.27 
    페이지 : 145 ~ 207


## 스프링의 두 번째 가치, 테스트

원래 절로 뺼 필요는 없지만, 주제가 "테스트"인 만큼 책에서 설명하는 가치를 짧게나마 짚고 가는게 좋을 것 같아서 따로 빼두었다. 저자 토비님은 스프링이 개발작에게 제공하는 가장 중요한 가치 2가지를 다음과 같이 뽑았다.

1. 개발자가 IoC/DI를 이용해서 손쉽게 OOP를 적용하게끔 도와주는 것.
2. 테스트

테스트는 만들어진 코드를 확신하게 해주며 변화에 유연하게 대처할 수 있게 해준다. 이제 스프링 프레임워크에서 어떻게 테스트 코드를 적용할 수 있는지 살펴보자.


## UserDaoTest의 문제점

`UserDaoTest`의 문제점을 돌아보기 전에, 먼저 테스트에 대해서 간단히 살펴보자. `UserDao`같이 우리가 만든 객체들을 스프링 애플리케이션 속에서 어떻게 테스트할 수 있을까? 다음의 2가지 방법이 있다.

1. 최소한의 코드로(컨트롤러, 모델 etc) 빠르게 애플리케이션을 만들어 띄워서 기능을 테스트한다.
2. 단위 테스트를 한다.

먼저 첫 번째 방식은 보통 QA팀에서 솔루션을 테스트하는 방식이다. 애플리케이션을 실행시킨 후, 배포한 기능을 이리 저리 테스트해서 정상적으로 동작하는지 확인한다. 이 경우, 오래 걸릴뿐더러 추가 기능 외에 여러 코드를 작성해야 하는 불편함이 있다. 성공적으로 돌아가면 다행이지만, 애플리케이션이 비정상적인 종료를 하면 정말 지옥이다. 추가된 기능이 문제인지 그 외 애플리케이션 내 문제인지 확신할 수 없다. 하지만, 솔루션이 전체적으로 동작하는지 확인할 수 있기 때문에 상용 배포 이전에 반드시 거쳐야 할 테스트이기도 하다.

두 번째 방식은 이전 장에서 우리가 했듯이 관심사들만 작은 단위로 빠르게 테스트하는 것이다. 이런 테스트를 "단위 테스트"라고 부르기도 한다. 간혹, 여러 컴포넌트의 작용하는 것을 테스트해야 할 때도 있는데, 이 경우 "통합 테스트"라는 테스트 방식을 선택할 수 있다. 단위 테스트이건 통합 테스트이건 확실히 첫 번째 방식보다 빠르게 이루어지며 대부분 코드로써 자동화가 가능하다. 최근이라고 하긴 애매하지만, 테스트 코드를 작성하고 개발하는 것은 개발자들 사이에서 필수 요소로 손꼽히고 있다. 

뭐 단위 테스트에 대해 찬양하듯 써놨지만, 개인적으로 상용 제품을 내야 하기 위해서는 둘 다 반드시 필요한 작업이라고 생각한다. 두 번째 방식은 개발자의 몫이겠지만.. 이제 이전 `UserDaoTest`의 코드를 다시 한 번 살펴보자.

```java
class UserDaoTest {
    public static void main(String[] args) throws SQLException {
        ApplicationContext applicationContext = new AnnotationConfigApplicationContext(DaoFactory.class);
        UserDao dao = applicationContext.getBean(UserDao.class);
        String id = "gurumee";
        String name = "hyunwoo";
        String password = "ilovespring";

        User user = new User(id, name, password);
        dao.add(user);
        System.out.println(user.getId() + " register success");

        User user2 = dao.get(id);
        System.out.println(user2.getId() + " " + user2.getPassword());

        System.out.println(id + " " + user2.getId());
        System.out.println(name + " " + user2.getName());
        System.out.println(password + " " + user2.getPassword());
    }

    // ...
}
```

이 테스트 코드는 단위 테스트라고 부르기가 살짝 부족하다. 무엇이 부족할까? 책에서는 다음이 부족하다고 설명하고 있다.

1. 테스트를 위해서 매번 main을 직접 실행해야 한다.
2. `UserDao`의 기능이 추가되면 main에서 모든 메서드를 테스트해야 하든가, 여러 `UserDaoTest` 클래스를 만들어서 각각 테스트해야 한다.
3. 결과를 눈으로 확인해서 정상 값이 들어오는지 확인해야 한다.
4. 테스트 후 데이터베이스에 저장된 정보를 지워야, 다음 테스트가 성공한다.

이 문제들을 자바의 단위 테스트 프레임워크인 `JUnit`을 이용하여 고쳐보자.


## JUnit으로 넘어가기

`JUnit`은 자바의 단위 테스트 프레임워크이다. 이 프레임워크가 실행하는 테스트 코드를 만들기 위해서는 다음의 2가지를 따라야 한다.

1. 메소드 레벨이 `public`이어야 한다.
2. 위의 `@Test` 애노테이션을 붙여야 한다.

이를 토대로 `UserDaoTest`의 메인 메소드를 바꿔보자.

```java
class UserDaoTest {
    @Test
    @DisplayName("UserDao add and get test")
    public void test01() throws SQLException {
        ApplicationContext applicationContext = new AnnotationConfigApplicationContext(DaoFactory.class);
        UserDao dao = applicationContext.getBean(UserDao.class);
        String id = "gurumee";
        String name = "hyunwoo";
        String password = "ilovespring";

        User user = new User(id, name, password);
        dao.add(user);
        System.out.println(user.getId() + " register success");

        User user2 = dao.get(id);
        System.out.println(user2.getId() + " " + user2.getPassword());

        System.out.println(id + " " + user2.getId());
        System.out.println(name + " " + user2.getName());
        System.out.println(password + " " + user2.getPassword());
    }

    // ...
}
```


> 참고!
> 
> @DisplayName 애노테이션은 JUnit5의 기능입니다. 테스트 메소드의 기능을 문자열로 표시할 수 있습니다. JUnit5 이전에는 테스트 메소드 명으로 어떤 테스트인지 명시해주는 것이 관례입니다.


이렇게 하면, 빌드 툴로 테스트 코드를 간단히 실행할 수 있다. 터미널에는 다음을 입력하면 테스트를 실행할 수 있다. "main을 직접 실행해야 한다"라는 단점이 사라진다.

```bash
# 그래들의 경우
$ gradle test

# 메이븐의 경우
$ mvn test
```

`IDE`의 경우, 간단히 실행시키는 방법이 각각 있을 것이다. 이제 "결과를 눈으로 확인해서 정상 값이 들어오는지 확인해야 한다"라는 단점을 고쳐보자. 코드를 다음과 같이 변경한다.

```java
class UserDaoTest {
    @Test
    @DisplayName("UserDao add and get test")
    public void test01() throws SQLException {
        ApplicationContext applicationContext = new AnnotationConfigApplicationContext(DaoFactory.class);
        UserDao dao = applicationContext.getBean(UserDao.class);
        String id = "gurumee";
        String name = "hyunwoo";
        String password = "ilovespring";

        User user = new User(id, name, password);
        dao.add(user);
        System.out.println(user.getId() + " register success");

        User user2 = dao.get(id);
        System.out.println(user2.getId() + " " + user2.getPassword());

        assertEquals(id, user2.getId());
        assertEquals(name, user2.getName());
        assertEquals(password, user2.getPassword());
    }

    // ...
}
```

`assertEquals`는 객체의 동등성을 따져서 같으면 테스트 성공, 다르면 테스트 실패를 나타낸다. 이 때 보통의 IDE는 성공은 초록색 막대가, 실패는 빨간색 막대가 보이게 된다. 테스트 결과에 대해서 값을 일일이 확인하지 않아도 알 수 있다는 것이다.

그리고 "`UserDao`의 기능이 추가되면 main에서 모든 메서드를 테스트해야 하든가, 여러 `UserDaoTest` 클래스를 만들어서 각각 테스트해야 한다." 단점이 사라지는 것을 확인해보자. 실제 내 `UserDaoTest`의 다른 테스트 코드 부분이다. 

```java
class UserDaoTest {
    // ...
    @Test
    @DisplayName("동등성 테스트 - DaoFactory")
    public void test02() {
        DaoFactory factory = new DaoFactory();
        UserDao dao1 = factory.userDao();
        UserDao dao2 = factory.userDao();
        assertNotSame(dao1, dao2);
    }

    @Test
    @DisplayName("동등성 테스트 - ApplicationContext")
    public void test03() {
        ApplicationContext applicationContext = new AnnotationConfigApplicationContext(DaoFactory.class);
        UserDao dao1 = applicationContext.getBean(UserDao.class);
        UserDao dao2 = applicationContext.getBean(UserDao.class);
        assertSame(dao1, dao2);
    }

    @Test
    @DisplayName("xmlContext 테스트")
    public void test04() {
        ApplicationContext applicationContext = new GenericXmlApplicationContext("applicationContext.xml");
        UserDao dao1 = applicationContext.getBean(UserDao.class);
        UserDao dao2 = applicationContext.getBean(UserDao.class);
        System.out.println(dao1);
        assertSame(dao1, dao2);
    }
}
```

`UserDao`에 대한 테스트보다, `ApplicationContext`에서 `UserDao`가 싱글톤으로 나오는지 이런 류의 테스트이긴 하지만, 다른 main을 작성하든가 main 하나에서 다 확인할 필요가 없어짐을 확인할 수 있다. 이제 "테스트 후 데이터베이스에 저장된 정보를 지워야, 다음 테스트가 성공한다."라는 단점을 지우기 위해서 `UserDao`에 메소드를 늘려보자.

```java
package com.gurumee.chonangam.user.dao;

import com.gurumee.chonangam.user.domain.User;

import javax.sql.DataSource;
import java.sql.*;

public class UserDao {
    // ...

    public void deleteAll() throws SQLException {
        Connection c = dataSource.getConnection();

        PreparedStatement ps = c.prepareStatement("delete from users");
        ps.executeUpdate();

        ps.close();
        c.close();
    }

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

뭐 필요한 것은 `deleteAll`이긴 한데, 다음 장을 위해서 `getCount`도 추가한다. JUnit에는 각 테스트 메소드를 실행할 때 마다 필요한 객체를 미리 셋업해두던가, 아니면 필요 자원을 해제시킬 수가 있다. `UserDaoTest`를 다음과 같이 수정해보자.

```java
class UserDaoTest {
    private UserDao userDao;

    @BeforeEach
    public void setUp() throws SQLException {
        ApplicationContext applicationContext = new GenericXmlApplicationContext("applicationContext.xml");
        userDao = applicationContext.getBean(UserDao.class);
        User user = new User("test", "test", "test");
        userDao.add(user);
    }

    @AfterEach
    public void tearDown() throws SQLException {
        userDao.deleteAll();
    }

    @Test
    @DisplayName("UserDao get success test")
    public void test01() throws SQLException {
        String value = "test";
        User user = userDao.get(value);
        assertEquals(value, user.getName());
        assertEquals(value, user.getPassword());
    }

    @Test
    @DisplayName("UserDao get failed test")
    public void test02() throws SQLException {
        String value = "test2";
        Assertions.assertThrows(SQLException.class, () -> {
            userDao.get(value);
        });
    }

    @Test
    @DisplayName("UserDao add success test")
    public void test03() throws SQLException {
        User user = new User("test", "test", "test");

        Assertions.assertThrows(SQLException.class, () -> {
            userDao.add(user);
        });

        int count = userDao.getCount();
        assertEquals(1, count);
    }

    @Test
    @DisplayName("UserDao add failed test")
    public void test04() throws SQLException {
        String value = "test2";
        Assertions.assertThrows(SQLException.class, () -> {
            userDao.get(value);
        });
    }
}
```

add, get, 테스트를 각각 분리하였다. 또한 setUp, tearDown을 작성하였는데 이 두 메소드가 테스트 메소드 실행 시, `UserDao`를 생성하고 종료 시에 데이터베이스에서 "users"에 저장된 정보를 모두 삭제한다. `JUnit`으로 변경하였고 이 자체로 훌륭하지만 아직 개선점은 남아있다.

먼저, 데이터베이스를 실제 애플리케이션 운영할 때 참조하는 것과, 테스트 코드를 실행할 때 참조하는 것을 분리해야 한다. 왜냐하면 운영/테스트가 동일한 데이터베이스를 참조하면, 테스트 코드를 진행했을 때 운영 데이터가 지워지거나 수정이 되는 등 심각한 문제를 발생시킨다. 또한, 스프링의 DI(IoC 컨테이너)를 활용하고 있지 않고 있다. 이제 스프링이 지원하는 테스트 기능으로 코드를 개선시켜보자.


## 왜 스프링인가? 테스트!!

먼저 Spring에서 지원하는 테스트 기능으로써 가장 강력한 것은 DI를 지원하게 해준다는 것이다. 코드를 다음과 같이 변경한다.

```java
@SpringBootTest
class UserDaoTest {
    @Autowired
    private UserDao userDao;

    // ...
}
```

물론 이게 베스트 프랙티스는 아니다. 데이터 레이어 테스트는 또 따로 해두는 것이 좋다. 여기서는 "DaoFactory를 통해 등록한 UserDao 빈을 주입시킬 수 있다"라는 것에 초점을 맞추면 된다. `@SpringBootTest`는 테스트 코드를 진행할 때, 스프링 부트용 테스트 러너를 실행시킨다. 그리고 `@Autowired`는 필드에 붙이면 필드 인젝션이 된다. 따라서 테스트 코드 실행 시 스프링 DI 컨테이너가 우리가 등록한 `UserDao`를 주입한다.

이제, 데이터베이스를 분리해보자. 현재 데이터베이스 정보를 `DaoFactory`에서 먼저 확인해보자.

```java
@Configuration
public class DaoFactory {
    @Bean
    public UserDao userDao() {
        UserDao userDao = new UserDao(dataSource());
        return userDao;
    }

    @Bean
    public DataSource dataSource() {
        SimpleDriverDataSource dataSource = new SimpleDriverDataSource();
        dataSource.setDriverClass(com.mysql.jdbc.Driver.class);
        dataSource.setUrl("jdbc:mysql://localhost/springbook");
        dataSource.setUsername("spring");
        dataSource.setPassword("book");
        return dataSource;
    }
}
```

데이터베이스를 "springbook"에서 "testdb"로 바꿔서 참조하게끔 해보겠다. 먼저, `test` 디렉토리 밑에 적당한 위치에 `TestDaoFactory`를 만든다.

```java
@TestConfiguration
public class TestDaoFactory {
    @Bean
    public UserDao testUserDao() {
        UserDao userDao = new UserDao(testDataSource());
        return userDao;
    }

    @Bean
    public DataSource testDataSource() {
        SimpleDriverDataSource dataSource = new SimpleDriverDataSource();
        dataSource.setDriverClass(com.mysql.jdbc.Driver.class);
        dataSource.setUrl("jdbc:mysql://localhost/testdb");
        dataSource.setUsername("spring");
        dataSource.setPassword("book");
        return dataSource;
    }
}
```

그리고, 이제 `UserDaoTest`를 다음과 같이 변경한다.

```java
@SpringBootTest
@Import(value = {TestDaoFactory.class})
class UserDaoTest {
    // ...
}
```

이제 테스트를 돌려보면, 동작하는 것을 확인할 수 있다. (**다만, mysql에서 데이터베이스를 만들고 유저에게 접근 권한을 주어야 한다.**) 우리는 스프링이 지원하는 테스트를 이용하여, 컨텍스트를 분리하였고 결과적으로 운영 환경에 영향을 주지 않고 테스트 코드를 수행할 수 있게 되었다. 또한 필드 인젝션을 통해서 빈이 주입받는 것 역시 알 수 있었다.

책에서는 `GenericXmlApplicationContext` 상에서 `test-applicationContext.xml` 파일을 만들어서, 테스트용 컨텍스트를 만드는 방법이 나온다. 위의 코드는 `AnnotationConfigApplicationContext`에서 테스트용 설정 빈인 `@TestConfiguration`을 이용하여 테스트 환경을 분리하는 방법이다. 그러나 이 방법들은 코드 상에 데이터베이스의 중요 정보들이 노출되기 때문에, 현재는 이 두 방법보단 `application.properties`이나 `application.yml` 설정 파일과 프로파일을 이용하여, 테스트 환경을 분리하거나 `testcontainers`를 이용하여 환경을 분리하는 방법을 쓴다.


## 스터디원들의 생각 공유

### 나의 질문과 답

1) Spock 이런 것도 있던데 JUnit과 무슨 차이가 있나?
2) 이런거 쓰는가?
3) 실제 테스트 코드 작성 비율은??
   
### 스터디원들의 질문과 답

- 불참

### 면접 질문으로 생각해볼 것?

- 불참