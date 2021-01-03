# 5장 서비스 추상화 (2)

![logo](./../logo.png)

    책장속 먼지털기 스터디 8차
    스터디 날짜 : 2021.01.04
    작성 날짜 : 2021.01.03 - 2021.01.04
    페이지 : 349 - 399


## 개요

이번 스터디에서 다룰 목표는 다음과 같다.

* UserService의 DB 트랜잭션 적용
* MailService 적용

각 기술들을 적용하면서 서비스 추상화를 적용하는 내용을 다룰 것이다. 또한 서비스 추상화를 하면서 예외 상황을 발생시키는 테스트를 작성하거나 Mock 객체를 이용한 테스트를 작성하는 방법에 대해서 다룰 것이다. 

## 트랜잭션 서비스 추상화
### 임의로 예외를 어떻게 발생시킬까?

책에서 요구 사항 한 개를 추가하였다. 유저들의 레벨을 업그레이드 하는 도중 어떤 이유로 서버가 다운 됐을 때, 이미 업그레이드 된 유저들의 상태를 돌려놓고, 레벨 변경 작업을 이루지 못했다고 유저들에게 알린다고 한다.

이를 테스트하려면 어떻게 할까? 어떻게 강제로 테스트 코드 내에서 에러를 발생시키는지 알아보자. 책에서는 `UserService`를 확장하여, 에러를 발생시키는 테스트 용 `UserService`를 만들 것을 추천하고 있다. 책에서는 `Setter Injection + xml` 기반으로 빈을 설정하기 때문에, 무리 없이 바로 진행할 수 있다. 하지만 나의 경우는 `Constructor Injection + Java` 기반으로 빈을 설정하고 있어서 다소 어려움이 있다. 나의 경우에는 합성 패턴으로 해결을 본다. 나의 `TestUserService`는 다음과 같다.

> 참고!
> 
> 책에서는 단순 UserService 상속, 필드 id에 대한 생성자만 존재합니다.

TestUserService.java
```java
public class TestUserService {
    private String id;
    private UserService userService;

    public TestUserService(String id, UserService userService) {
        this.id = id;
        this.userService = userService;
    }

    public boolean canUpgradeLevel(User user) {
        return userService.canUpgradeLevel(user);
    }

    public void upgradeLevel(User user) {
        if (user.getId().equals(this.id)) {
            throw new TestUserServiceException();
        }

        userService.upgradeLevels();
    }

    public void upgradeLevels() {
        List<User> users = userService.getUserDao().getAll();

        for (User user : users) {
            if (canUpgradeLevel(user)) {
                upgradeLevel(user);
            }
        }
    }

    public void add(User user) {
       userService.add(user);
    }
}
```

책과 달리 `UserService.upgradeLevels` 변경될 때마다 이를 변경해주어야 한다는 단점이 생겼지만, 그런대로 쓸 수 있을 것 같다. 또한 강제 예외를 발생시키기 위한 `TestUserServiceException`을 생성한다.

TestUserServiceException.java
```java
public class TestUserServiceException extends RuntimeException {
}
```

이제 `UserServiceTest.java`를 다음과 같이 테스트 코드를 추가한다.

UserServiceTest.java
```java
// ...
@Test
@DisplayName("예외 발생 시 작업 취소 여부 테스트")
public void test_cancel_when_exception() {
    TestUserService testUserService = new TestUserService(users.get(3).getId(), userService);

    assertThrows(TestUserServiceException.class, () -> {
        testUserService.upgradeLevels();
    });

    checkLevel(users.get(1), false);
}
// ...
```

테스트 코드를 돌려보면, 다음과 같이 "테스트 실패"라는 결과를 얻을 수 있다.

```
expected: <BASIC> but was: <SILVER>
org.opentest4j.AssertionFailedError: expected: <BASIC> but was: <SILVER>
...
```

### 트랜잭션이 무엇인가?

위의 테스트 실패는 트랜잭션 문제이다. 트랜잭션이란, DB에서 "더 이상 나눌 수 없는 작업"을 뜻한다. 레벨을 업그레이드하는 로직을 다시 한 번 살펴보자.

UserService.java
```java
// ...
public void upgradeLevels() {
    List<User> users = userDao.getAll();

    for (User user : users) {
        if (canUpgradeLevel(user)) {
            upgradeLevel(user);
        }
    }
}

// ...
```

모든 유저를 순회해서, 업그레이드 가능한 유저는 레벨을 한 단계 업그레이드 시킨다. 만약 이 순회 도중 서버에 에러가 발생해서 기능이 다운된다면 어떻게 될까? 기본적으로 자바 애플리케이션에서는 순회된 유저까지만 업그레이드하고 멈출 것이다. 논리적으로는 업그레이드 도중 멈추기보다, 이전 작업 상태로 돌아가는 것이 맞을 것이다. 이렇게 **DB에 어떤 작업을 논리적으로 묶어 실패했을 때 어떻게 하는가를 정하는 것이 바로 트랜잭션**이다. 일종의 작업 단위의 경계라고 보면 된다.

일반적으로 트랜잭션을 시작하는 방법은 한 가지이지만, 끝내는 방법은 다음과 같이 2가지가 있다.

1. Commit
2. Rollback

첫 번째 `Commit`은 작업이 성공적으로 마무리되었다고 DB에 알려 내용을 수정하는 것이다. 두 번째 `Rollback`은 작업이 실패하였음을 알려 DB 내용을 작업 이전으로 되돌려 놓는다. 이제 `JDBC`에서 어떻게 트랜잭션을 적용하는지 알아보자. 아래는 간단한 트랜잭션 적용하는 코드의 예문이다.

```java
Connection c = dataSource.getConnection();
c.setAutoCommit(false); // 트랜잭션 시작.

try {
    // 트랜잭션 설정하고 싶은 SQL 구문들
    c.commit();     // 성공적이면 트랜잭션 커밋!
} catch(Exception e) {
    c.rollback();   // 도중 실패하면 롤백
} finally {
    c.close();      // 마무리는 Connection 리소스 해제
}
```

이제 본격적으로 트랜잭션이 무엇이고 `JDBC`에서 어떻게 적용하는지 알았으니 `UserService`에 트랜잭션을 적용해보자. 

### UserService에 트랜잭션 적용

1. UserService에 트랜잭션 적용
2. DB 기술 독립적인 트랜잭션의 적용

## 메일 서비스 추상화

1. JavaMailService
2. 테스트를 위한 서비스 추상화
3. Mock 객체 생성 및 테스트