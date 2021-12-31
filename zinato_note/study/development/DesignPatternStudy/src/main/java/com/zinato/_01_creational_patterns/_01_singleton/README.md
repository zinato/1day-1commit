# Singleton

> 오직 한개의 클래스 인스턴스만을 갖도록 보장하고, 이에 대한 전역적인 접근점(access point)을 제공 
- 시스템 런타임, 환경 셋팅에 대한 정보등, 인스턴스가 여러개 일 때 문제가 생길 수 있는 경우가 있음.
***인스턴스를 오직 한개만*** 만들어 제공하는 클래스가 필요할 때 사용 한다. 
  - 게임 설정 화면
  - JDBC Connection 

### 기본적인 Singleton 패턴 

- Singleton 패턴은 인스턴스 한개만 제공해야 하기 때문에 외부에서 new 생성자로 생성하면 안됨 
  - 그렇기 때문에 생성자를 private 으로 만들어 줌 
- global 에서 접근할 수 있는 접근점을 static getInstance()로 제공
```java
public class Settings {

  private static Settings instance;
  //New 키워드로 생성할 수 없음
  private Settings() {}

  public static Settings getInstance() {
    if (instance == null) {
      instance = new Settings();
    }
    return instance;
  }
}
```

- 문제점 
  - 위 코드는 ***멀티 쓰레드*** 환경에서 안전한가? 

### 멀티 쓰레드 환경에서 안전하게 구현하는 방법

#### 1. synchronized 키워드 사용 
```
public static synchronized  Settings getInstance() {
    if (instance == null) {
      instance = new Settings();
    }
    return instance;
  }
```

- synchronized 키워드로 생성하게 되면 멀티 쓰레드 상황에서 안전해진다. 
- 하지만 getInstance()를 동기화 하는데 비용이 많이 들어 성능상 단점이 존재한다.

#### 2. 이른 초기화 (eager initialization) 사용하기 

```java
public class Settings {

  private static final Settings INSTANCE = new Settings();
  private Settings() {}

  public static synchronized  Settings getInstance() {
    return INSTANCE;
  }
}
```

- 이 방법 역시 Thread Safe
  - 여러 쓰레드가 들어오더라도 이미 생성한 INSTANCE를 return 해주기 때문
  - INTANCE는 Settings Class가 로딩되는 시점에 static field들이 초기화 되기 때문
- 단점 
  - 처음 초기화 하는 작업이 오래 걸리면 성능이 느려짐
  - 해당 객체를 사용하지 않으면 사용하지도 않는 객체를 미리 만들어두는 것이 됨(심지어 초기화도 느리면 ...)

#### 3. double checked locking 사용하기

```java
public class Settings {

  private static volatile Settings instance;
  //New 키워드로 생성할 수 없음
  private Settings() {}

  public static Settings getInstance() {
    if (instance == null) {
      synchronized (Settings.class) {
        if (instance == null) {
          instance = new Settings();
        }
      }
    }
    return instance;
  }
}
```

- check를 한번 하고 synchronized block 에서 한번 더 check 한다고 하여 double checked locking 이라고 함
- volatile 키워드를 사용해줘야 함 
- 매번 synchronized가 걸리지 않아 위에 나온 synchronized 사용한 패턴보다 성능에서 유리한 이점 

- 단점
  - 복잡한 패턴임 
  - volatile을 왜 쓰는지 이해도가 필요 
  - 멀티 쓰레드 환경에서 메모리를 관리하는 방법을 이해해야 제대로 사용할 수 있음 
  - jdk 1.5 버전 이상에서만 동작함 

#### 4. static inner 클래스 사용하기 (권장)

```java
public class Settings {

  private Settings() {}

  private static class SettingsHolder {
    private static final Settings INSTANCE = new Settings();
  }
  public static Settings getInstance() {
    return SettingsHolder.INSTANCE;
  }
}
```

- 멀티 쓰레드 환경에서도 안전 
- getInstance() 가 호출될 때 SettingsHolder.INSTANCE가 만들어 짐 (Lazy Loading 가능)

### 싱글톤 패턴 구현 방법을 깨트리는 방법

#### Reflection 사용하기 
 
```java
public class App {
  public static void main(String[] args) throws NoSuchMethodException, InvocationTargetException, InstantiationException, IllegalAccessException {
    Settings settings1 = Settings.getInstance();
    
    Constructor<Settings> constructor = Settings.class.getDeclaredConstructor();
    constructor.setAccessible(true);
    Settings settings2 = constructor.newInstance();
    System.out.println(settings1 == settings2); //false
  }
}
```

- 사용자가 위와 같이 잘못 작성하면 싱글턴 패턴이 깨짐 
- 대응 방법이 딱히 없음 

#### 직렬화 & 역직렬화 사용하기 

```java
public class App {
  public static void main(String[] args) throws IOException, ClassNotFoundException {
    Settings settings1 = Settings.getInstance();
    Settings settings2 = null;

    try (ObjectOutput out = new ObjectOutputStream(new FileOutputStream("settings.obj"))){
      out.writeObject(settings1);
    }
    try (ObjectInput in = new ObjectInputStream(new FileInputStream("settings.obj"))) {
      settings2 = (Settings) in.readObject();
    }
    System.out.println(settings1 == settings2); //false
  }
}
```

- 역직렬화 대응 방안 
```java
public class Settings implements Serializable {

  private Settings() {}

  private static class SettingsHolder {
    private static final Settings INSTANCE = new Settings();
  }
  public static Settings getInstance() {
    return SettingsHolder.INSTANCE;
  }

  protected Object readResolve() { //역직렬화 할 때 여기를 사용 
    return getInstance();
  }
}
```

### 안전하고 단순하게 구현 하는 방법 

- Enum 으로도 간단하게 구현 가능 
- Constructor, Getter, Setter 모두 구현 가능 
- reflection에 안전한 코드 
- Enum Class 자체가 Serializable을 구현한 Class 

```java
public enum SettingsWithEnum {
  INSTANCE;
  //생성자 , Getter, Setter 모두 가능 ,
  // reflection에 안전한 코드
  SettingsWithEnum() {
  }
}
```

### 싱글톤 (Singleton) 패턴 복습 

- 자바에서 enum을 사용하지 않고 싱글톤 패턴을 구현하는 방법은?
- private 생성자와 static 메소드를 사용하는 방법의 단점은?
- enum 을 사용해 싱글톤 패턴을 구현하는 방법의 장점과 단점은 ?
  - 장점 : reflection에 안전, multi thread safe, 간단한 구조 
  - 단점 : 다른 class 상속을 쓰지 못함 , 오로지 enum만 상속 가능, **lazy loading**을 사용하고 싶으면 holder를 사용한 패턴을 사용해야함.  
- static inner 클래스를 사용해 싱글톤 패턴을 구현하라. 


### 실무에서는 ?

- 스프링에서 Bean의 스코프 중에 하나로 싱글톤 스코프.
- 자바 java.lang.Runtime 
- 다른 디자인 패턴(빌더, 퍼사드, 추상 팩토리 등) 구현체의 일부로 쓰이기도 함.  