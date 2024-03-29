# ITEM 04 : 인스턴스화를 막으려면 private 생성자를 사용하라

- 단순히 정적 메서드와 정적 필드만을 담은 클래스를 만들 경우가 있다. 
- 객체 지향적으로 보면 좋지 않지만 쓰임새가 있을 때가 있다. 
- java.lang.Math와 java.util.Arrays 처럼 기본 타입값이나 배열 관련 메서드를 모을 수 있다. 
- java.util.Collections 처럼 특정 인터페이스를 구현하는 개체를 생성해주는 정적 메서드를 모을 수 있다.
- __추상클래스를 만드는 것으로는 인스턴스화를 막을 수 없다.__
- 컴파일러가 기본 생성자를 만드는 경우는 오직 명시된 생성자가 없을 때 뿐인다. __private 생성자를 추가하면 클래스의 인스턴스화를 
막을 수 있다.__
  
```java
public class UtilityClass {
  //기본 생성자가 만들어지는 것을 막는다.
  private UtilityClass() {
    throw new AssertionError();
  } 
}
```
  
- 어떤 환경에서도 인스턴스화를 막아주지만 생성자가 있는데도 호출할수 없는 것이 직관적이지 않으므로 주석을 추가 해주자.
- 이 방식은 상속도 불가능하게 한다. 하위 클래스가 상위 클래스의 생성자에 접근할 길이 막히기 때문이다. 
