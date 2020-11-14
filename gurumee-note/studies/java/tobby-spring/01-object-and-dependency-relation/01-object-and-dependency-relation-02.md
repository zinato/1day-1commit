# 오브젝트와 의존관계 (2)

![logo](./../logo.png)

    책장속 먼지털기 스터디 2차
    스터디 날짜 : 2020.11.16 (불참)
    작성 날짜 : 2020.11.14 ~ 15
    페이지 : 102 ~ 143


## 싱글톤 레지스트리? 

우리가 만든 `DaoFactory`를 직접 사용하는 것과, `@Configuration` 애노테이션으로 `ApplicationContext`에 등록해서 사용하는 것에는 어떤 차이가 있을까? 결론부터 말하면 차이는 스프링 빈이냐 아니냐의 차이다. 스프링 빈은 이전에 언급했 듯이, 스프링 IoC 컨테이너가 라이프 사이클을 관리하는 객체들이다. 스프링 빈은 특정 스코프를 가지는데 기본적으로 "싱글톤" 스코프를 가진다. 쉽게 말하면, 애플리케이션에서 딱 1개의 객체만 만들어진다.

이 후 챕터에서 다룰 것이지만 한 발 앞서서 JUnit(나는 JUnit5이다.) 라이브러리를 활용하여, 테스트 코드를 추가해보자. 


```java
// ..

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
}
```

위의 2 테스트는 `DaoFactory`를 직접 사용하는 것과, 스프링 빈으로 등록시켜서 사용할 때의 차이를 명확하게 보여준다. `assertNotSame`은 두 객체의 동일성이 틀린지에 대해 테스트한다.(dao1 != dao2) `assertSame`은 두 객체의 동일성이 같은지에 대해 테스트한다.(dao1 == dao2) 즉, 직접 `DaoFactory`를 만들어서 사용할 때, `userDao` 메소드를 호출할 때마다 새로운 객체를 반환한다. 반면에, 스프링 빈으로 등록했을 경우, `DaoFactory`, `UserDao` 모두 빈으로 등록되기 때문에 같은 객체를 반환하게 된다.

> 참고! 동일성과 동등성
> 
> "오브젝트가 동일하다", "오브젝트가 동등하다"라는 말은 얼핏 들으면 같은 말로 이해할 수 있습니다. 그러나 이들은 서로 다른 말입니다. 동일하다는 같은 객체인지를 판단합니다. 즉, 객체의 참조가 같은지를 따집니다. 자바에서는 "==" 연산자로 이를 확인할 수 있습니다. 반면에, 동등성은 객체에 들어있는 정보가 같은지를 판단합니다. 자바에서는 보통 "equals" 메소드를 호출하여 동등성을 판단하는 것이 관례입니다. 다만, 클래스에 동등성을 판단하는 equals 메소드가 적절하게 오버라이딩 되어 있지 않으면, 동등성이 꺠질 수 있다는 것을 알아두세요.

그렇다면 여기서 의문점이 생긴다. "왜 스프링은 (기본적으로 )싱글톤으로 스프링 빈을 만드는 것인가?" 이는 스프링의 출발이, 자바 엔터프라이즈 기술을 사용하는 서버 환경에서 시작되었기 때문이다. 

스프링이 시작됐던 시기에, 서버 하나 당 최대 초당 수십/수백 번씩 브라우저나 다른 시스템으로부터 요청을 처리할 수 있는 성능이 요구되었다. 이 때 매번 요청 시마다, 그 요청을 처리하는 각 컴포넌트를 매번 생성했다고 가정해보자. 이를 감당할 수 있었겠는가? 요청 1번의 5개의 객체가 만들어지고, 초당 500개 요청이 들어온다고 하면, 초당 2500개의 객체가, 분당은 5만개, 한시간이면 9백만개가 생성된다. GC 성능이 좋아진 지금이라도 부하를 받을 수 밖에 없다. 그래서 이 자원들을 최대한 아끼기 위해서, 하나의 객체를 공유해서 동시에 사용하게 되었다. 이를 디자인 패턴으로 시각으로는 "싱글톤 패턴"이라고 한다. 이 패턴에 대한 자세한 내용은 다음 문서를 참고하자.

* [싱글톤 패턴](https://gmlwjd9405.github.io/2018/07/06/singleton-pattern.html)

그러나 "싱글톤 패턴"은 다음과 같은 한계가 있다.

1. private 생성자를 갖고 있기 때문에 상속할 수 없다.
2. 테스트하기 어렵다.
3. 서버 환경에서 하나만 만들어짐을 보장할 수 없다.
4. 싱글톤은 전역 상태로 만들 수 있기 떄문에 바람직하지 못하다. (정적 메소드로 접근할 수 있다.)

이러한 한계를 극복하기 위해서, 스프링은 "싱글톤 레지스트리"를 사용한다. 쉡게 말해서, 싱글톤 패턴이 적용되지 않은 일반 클래스조차도 싱글톤으로 관리할 수 있게 만드는 기술이다. `ApplicationContext`는 "빈 팩토리"이자, "싱글톤 레지스트리"이기도 하다. 그래서 이전 테스트 코드에서 확인할 수 있듯이 `getBean` 메소드로 스프링 빈을 수 없이 호출하더라도, 딱 하나의 객체만 반환됨을 확인할 수 있다. 물론 아무런 설정한 것이 없다면 말이다. 참고적으로 자주 사용되는 스프링 빈의 스코프는 다음과 같다.

1) 싱글톤 - 기본 스코프, 애플리케이션 전체에서 딱 1개
2) 프로토타입 - 빈을 호출할 때마다, 생성됨.
3) 리퀘스트 - HTTP 요청 시 빈이 생성됨.
4) 세션 - 웹의 세션과 스코프가 유사함.

이러니 저러니 해도, 스프링 빈의 스코프는 아마 싱글톤이 제일 많을 것이다. 왜냐하면 기본 스코프이기도 하고, 앞서 언급했듯 자원을 최소화할 수 있기 떄문이다. 이 때 주의 사항이 있다. 바로 싱글톤 객체가 "상태"를 가져서는 안된다는 것이다. 왜냐하면 상태를 가지게 된다면, 멀티 스레딩 환경에서, 그 상태가 의도치 않게 변할 수 있다. 잘 이해가 안된다면 다음 코드를 보자.

```java
public class UserDao {
    private ConnectionMaker connectionMaker;
    // 추가 필드
    private Connection c;

    public UserDao(ConnectionMaker connectionMaker) {
        this.connectionMaker = connectionMaker;
    }

    public void add(User user) throws ClassNotFoundException, SQLException {
        // 수정
        c = connectionMaker.makeConnection();

        PreparedStatement ps = c.prepareStatement("insert into users(id, name, password) values(?, ?, ?)");
        ps.setString(1, user.getId());
        ps.setString(2, user.getName());
        ps.setString(3, user.getPassword());
        ps.executeUpdate();

        ps.close();
        c.close();
    }

    // ...
}
```

`UserDao`를 살짝 변경한 것인데, 필드로 `Connection`을 가지게 되었다. 이 경우, `add` 메서드를 호출할 때 마다, c가 참조하는 `Connection` 객체가 변경된다. 이는 멀티 스레딩 환경에서 값이 제대로 데이터베이스에 저장되지 않는 심각한 문제를 초래한다. 또한, 이 메소드의 실행 결과가 잘되는 것조차 알 수 없다. 이러한 문제를 제거하기 위해서는 이전 코드처럼 상태가 변경될 수 있는 객체는 메소드 레벨에서 관리하는 것이 좋다. 혹은 필드로 주입하는 객체를 빈으로 만드는 방법도 있지만.. 별로 좋은 생각은 아니라고 본다.


## DI

## Xml 기반 설정


## 스터디원들의 생각 공유

### 나의 질문과 답

1) 
2) 


### 스터디원들의 질문과 답

- 불참

### 면접 질문으로 생각해볼 것?

- 불참