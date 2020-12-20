# 5장 서비스 추상화 (1)

![logo](./../logo.png)

    책장속 먼지털기 스터디 7차
    스터디 날짜 : 2020.12.21
    작성 날짜 : 2020.12.20
    페이지 : 317 - 348


## 개요

책에서 나오는 "서비스 추상화"란 성격이 비슷한 여러 종류의 기술을 추샇와하고 이를 일관된 방법으로 사용할 수 있도록 하는 것이다. 5장 전체에서는 `DAO`에 트랜잭션을 적용하면서 어떻게 서비스를 추상화할 수 있는지 알아보지만, 이번 주차 스터디 범위에서는 그 이전 단계까지의 코드를 작성한다.


## 사용자 레벨 관리 기능 추가

### 기능을 위한 필드 추가 및 테스트 코드 수정

현재 `UserDao`는 "CRUD" 기능 외에 어떠한 기능도 가지고 있지 않다. 여기에 다음 비지니스 로직을 추가한다.

* 사용자 레벨은 BASIC, SILVER, GOLD 중 하나이다.
* 처음 가입하면 BASIC이다.
* 가입 후 50회 이상 로그인을 하면 BASIC -> SILVER
* SILVER 레벨이면서 추천 수 30회 이상이면 SILVER -> GOLD
* 사용자 레벨의 변경 작업은 일정한 주기를 가지고 일괄적으로 진행된다.

이제 `User` 객체에 필드를 추가한다. 

* 사용자 레벨 Enum Level 
* 로그인 횟수 login
* 추천 횟수 recommend

해당 테이블을 다음과 같이 변경하도록 한다.

```sql
drop table if exists users;
create table users (
    id varchar(10) primary key,
    name varchar(20) not null,
    password varchar(20) not null,
    level int not null,
    login int not null,
    recommend int not null
);


create database testdb;
use testdb;
drop table if exists users;
grant all privileges on testdb.* to spring@'%';


create table users (
    id varchar(10) primary key,
    name varchar(20) not null,
    password varchar(20) not null,
    level int not null,
    login int not null,
    recommend int not null
);
```

이제 코드를 수정해보자. 먼저 레벨을 표현하는 Enum 클래스이다.

Level.java
```java
@Getter
public enum Level {
    BASIC(1), SILVER(2), GOLD(3);

    private final int value;

    Level(int value) {
        this.value = value;
    }

    public static Level valueOf(int value) {
        switch (value) {
            case 1: return BASIC;
            case 2: return SILVER;
            case 3: return GOLD;
            default: throw new AssertionError("Unknown value: " + value);
        }
    }
}
```

`User`를 다음과 같이 변경한다.

```java
@NoArgsConstructor @AllArgsConstructor
@Getter @Setter @ToString @EqualsAndHashCode
// 추가 애노테이션
@Builder
public class User {
    private String id;
    private String name;
    private String password;

    // 추가 필드
    private Level level;
    private int login;
    private int recommend;
}
```

이제 `UserDaoTest`를 수정하자.

```java
@SpringBootTest
@Import(value = {TestDaoFactory.class})
class UserDaoTest {
    @Autowired
    private UserDao userDao;

    private User user;

    @BeforeEach
    public void setUp() {
        user = User.builder()
                .id("test1")
                .name("test1")
                .password("test1")
                .level(Level.BASIC)
                .login(1)
                .recommend(0)
                .build();
        userDao.add(user);
        User tmp = User.builder()
                .id("test2")
                .name("test2")
                .password("test2")
                .level(Level.SILVER)
                .login(55)
                .recommend(10)
                .build();
        userDao.add(tmp);
        tmp = User.builder()
                .id("test3")
                .name("test3")
                .password("test3")
                .level(Level.GOLD)
                .login(100)
                .recommend(40)
                .build();
        userDao.add(tmp);
    }

    @AfterEach
    public void tearDown() {
        userDao.deleteAll();
    }

    @Test
    @DisplayName("UserDao get success test")
    public void test01() {
        User findUser = userDao.get("test1");
        assertEquals(user.getName(), findUser.getName());
        assertEquals(user.getPassword(), findUser.getPassword());
        assertEquals(user.getLevel(), findUser.getLevel());
        assertEquals(user.getLogin(), findUser.getLogin());
        assertEquals(user.getRecommend(), findUser.getRecommend());
    }

    @Test
    @DisplayName("UserDao get failed test")
    public void test02() {
        String value = "test4";
        Assertions.assertThrows(EmptyResultDataAccessException.class, () -> userDao.get(value));
    }

    @Test
    @DisplayName("UserDao add success test")
    public void test03() {
        User user = User.builder()
                .id("test4")
                .name("test4")
                .password("test4")
                .level(Level.BASIC)
                .login(0)
                .recommend(0)
                .build();
        userDao.add(user);
        int count = userDao.getCount();
        assertEquals(4, count);
    }

    @Test
    @DisplayName("UserDao add failed test")
    public void test04() {
        User user = User.builder()
                .id("test1")
                .name("test1")
                .password("test1")
                .level(Level.BASIC)
                .login(0)
                .recommend(0)
                .build();
        Assertions.assertThrows(DuplicateKeyException.class, () -> userDao.add(user));

        int count = userDao.getCount();
        assertEquals(3, count);
    }

    @Test
    @DisplayName("UserDao getAll test")
    public void test05() {
        List<User> list = userDao.getAll();
        assertEquals(3, list.size());

        User user = list.get(0);
        assertEquals("test1", user.getId());
        assertEquals("test1", user.getName());
        assertEquals("test1", user.getPassword());

        for (int i=4; i<=8; i++) {
            String msg = "test" + i;
            User tmp = User.builder()
                    .id(msg)
                    .name(msg)
                    .password(msg)
                    .level(Level.BASIC)
                    .login(0)
                    .recommend(0)
                    .build();
            userDao.add(tmp);
        }

        list = userDao.getAll();
        assertEquals(8, list.size());

        for (int i=1; i<5; i++) {
            String expected = "test" + (i+1);
            User tmp = list.get(i);
            assertEquals(expected, tmp.getId());
            assertEquals(expected, tmp.getName());
            assertEquals(expected, tmp.getPassword());
        }
    }
}
```

다른 부분은 `User`의 필드가 바뀌는 것에 대해 변경된 테스트 코드이다. 중요한 부분은 픽스처가 추가되고 빌더로 추가된 필드까지 모두 저장하는 부분이다.

```java
private User user;

@BeforeEach
public void setUp() {
    //픽스처 설정
    User user = User.builder()
            .id("test1")
            .name("test1")
            .password("test1")
            .level(Level.BASIC)
            .login(1)
            .recommend(0)
            .build();
    userDao.add(user);

    다른 데이터 저장
    user = User.builder()
            .id("test2")
            .name("test2")
            .password("test2")
            .level(Level.SILVER)
            .login(55)
            .recommend(10)
            .build();
    userDao.add(user);
    user = User.builder()
            .id("test3")
            .name("test3")
            .password("test3")
            .level(Level.GOLD)
            .login(100)
            .recommend(40)
            .build();
    userDao.add(user);
}
```

이제 테스트 `getTest`를 한 번 다음과 같이 수정해보자.

```java
@Test
@DisplayName("UserDao get success test")
public void test01() {
    User findUser = userDao.get("test1");
    assertEquals(user.getName(), findUser.getName());
    assertEquals(user.getPassword(), findUser.getPassword());
    assertEquals(user.getLevel(), findUser.getLevel());
    assertEquals(user.getLogin(), findUser.getLogin());
    assertEquals(user.getRecommend(), findUser.getRecommend());
}
```

실패한다. 왜 실패할까. 현재 `UserDaoJdbc`는 추가된 필드를 파싱해서 값을 불러오는 코드와 저장하는 코드가 없기 때문이다. `UserDaoJdbc`를 다음과 같이 변경한다.

```java
@NoArgsConstructor @AllArgsConstructor
@Getter @Setter
public class UserDaoJdbc implements UserDao {
    private JdbcTemplate jdbcTemplate;
    private final RowMapper<User> rowMapper = (rs, rowNum) -> {
        User user = new User();
        user.setId(rs.getString("id"));
        user.setName(rs.getString("name"));
        user.setPassword(rs.getString("password"));
        user.setLevel(Level.valueOf(rs.getInt("level")));
        user.setLogin(rs.getInt("login"));
        user.setRecommend(rs.getInt("recommend"));
        return user;
    };

    public void add(User user) throws DataAccessException {
        String query = "insert into users(id, name, password, level, login, recommend) values(?, ?, ?, ?, ?, ?)";
        jdbcTemplate.update(query, user.getId(), user.getName(), user.getPassword(), user.getLevel().getValue(), user.getLogin(), user.getRecommend());
    }

    // ...
}
```

자 이제 테스트를 돌려보면 잘 통과하는 것을 확인할 수 있다.

### 수정 기능 추가

### 서비스 코드의 등장!

### 리팩토링 리팩토링!