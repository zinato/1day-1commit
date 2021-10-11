# ITEM 03. private 생성자나 열거 타입으로 싱글턴임을 보증하라

- 싱글턴(singleton)이란 인스턴스를 오직 하나만 생성할 수 있는 클래스를 말한다.[Gamma95]
- 클래스를 싱글턴으로 만들면 이를 사용하는 클라이언트를 테스트하기 어려워질 수 있다. 
- 싱글턴을 만드는 방식 
    -  private으로 생성자를 감춘다. 
        1. 인스턴스에 접근할 수 있는 유일한 public static final 멤버를 만든다.
            ```java
           public class Elvis {
           
            // 외부에서 접근할 수 있는 유일한 수단 (final)로 선언되어 있음
            // public static 필드가 final이라 절대로 다른 객체를 참조할 수 없다
            public static final Elvis INSTANCE = new Elvis();
            // 생성자를 private 으로 감춘다.
            private Elvis() { ... }
            
            public void leaveTheBuilder() { ... }
           }
           ```
           - private 생성자는 public static final 필드인 Elvis.INSTANCE를 초기화활 때 딱 한번 호출.
           - public, protected 생성자가 없으므로 인스턴스가 전체 시스템에서 하나만 보장.
           - 예외는 하나, 리플렉션 API인 AccessibleObject.setAccessible을 사용해 private 생성자를 호출할 수 있다. 이 때 
            생성자를 통해 두번째 객체가 생성하려고 할 때 예외를 던지게 하면 된다.
           - API에 해당 클래스가 싱글턴임이 명백히 드러난다.
           - 간결하다. 
        2. 정적 팩토리 메서드를 public static 멤버로 제공한다.
            ```java
            public class Elvis {
              private static final Elvis INSTANCE = new Elvis();
              //private으로 생성자를 감춤 
              private Elvis() {}
              
              public static Elvis getInstance() {return INSTANCE;}
           
              public void leaveTheBuilder() {}
           }
           ```
       
           - Elvis.getInstance는 항상 같은 객체를 참조 반환 하므로 두번째 Elvis 인스턴스는 없다(리플렉션 예외 동일)
           - API를 바꾸지 않고도 싱글턴이 아니게 변경 가능.
           - 유일한 인스턴스를 반환하던 팩터리 메서드가 호출하는 스레드별로 다른 인스턴스를 넘겨주게 할 수 있다.
           - 정적 팩터리를 제니릭 싱글턴 팩터리 메서드로 만들 수 있다. 
           - 정적 팩토리의 메서드 참조를 공급자(supplier)로 할 수 있다. 
            - Elvis::getInstance를 Supplier<Elivs>로 사용 
    
        3. 싱글턴 클래스를 직렬화 하려면 단순히 Serializable을 구현으로는 부족하다. 
            - 모든 인스턴스 필드를 transient라고 선언하고 readResolve 메서드를 제공해야 한다. 
            - 이렇게 하지 않으면 인스턴스를 역직렬화 할 때 세로운 인스턴스가 만들어진다. 
            
            ```java
           //싱글턴 임을 보장해주는 readResolve 메서드 
               private Object readResolve() {
                //진짜 Elvis를 반환, 가짜 Elvissms 가비지 컬렉터에 맡긴다. 
                return INSTANCE; 
               }
           ```
       
        4. 싱글턴을 만드는 세번째 방법은 열거 타입을 선언 
        
            ```java
           public enum Elvis {
            INSTANCE;
           
            public void leaveTheBuilder() {}
           }
           ```
       
            - public 필드와 비슷하지만 더 간결하게 직렬화 가능. 
            - 아주 복잡하고 리플렉션 공격에도 제 2의 인스턴스가 생기는 일을 완벽하게 막아줌.
            - __대부분 상황에서 싱글턴을 만드는 가장 좋은 방법.__
            - 만들려는 싱글턴이 Enum 외의 값을 상속해야 한다면 사용할 수 없다. (열거타입이 다른 인터페이스를 구현하도록 선언할 수는 있다.)
    
