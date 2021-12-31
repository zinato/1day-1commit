# ITEM 06 : 불필요한 객체 생성을 피하라

```java
  //절대 사용하지 말것. 실행될 때마다 String 인스터스를 새로 만든다.
  String s = new String("bikini");

  //위와 완전히 똑같은 기능을 하지만 성능은 완전 다르다. 
  // 같은 가상 머신 안에서 똑같은 문자열 리터럴을 사용하는 모든 코드가 같은 객체를 재사용함이 보장된다. 
  String s = "bikini";
```

- Boolean(String) 생성자 대신 Boolean.valueOf(String) 팩토리 메서드를 사용하는 것이 좋다.

- String.matches 와 같이 비용이 많이 드는 메서드는 객체가 생성될 때 매번 생성하게 하는 것 보다 클래스 초기화 과정에서 직접 생성해 키생해두고 나중에 메서드가 호출될 때 이 인스턴스를 재사용하는 것이 성능상으로 유리하다.

```java
//Not Use
static boolean isRomanNumeral(String s) {
  return s.matches("^(?=.)$"); //String matches는 비용이 높다. 매 객체마다 생성해서 사용하면 속도 저하가 생길 수 있다. 
}

//Use
  public class RomanNumerals {
    private static final Pattern ROMAN = Pattern.compile("^(?=.)$");

    static boolean is RomanNumeral(String s) {
      return ROMAN.matcher(s).mathces();
    }
  }
```

- 박싱된 기본 타입보다는 기본타입을 사용하고 의도치 않은 오토박싱이 숨어들지 않도록 주의하자. 부주의하게 사용한 오토박싱으로 인해 어마어마한 속도차이가 발생할 수 있다.

- 이번 아이템을 "객체 생성은 비싸니 피해야 한다" 로 오해하면 안된다.

- 요즘 JVM은 작은 객체를 생성하고 삭제하는 것은 그렇게 크게 부담되는 작업은 아니다. 또 잘못 이해하여 객체 생성을 피하고자 개별 객체 풀(Pool)을 만들지 말자. 더 메모리 사용량을 늘리고 성능을 떨어뜨린다.(DB 객체 풀과 같은 매우 비싼 객체 풀은 제외)

- 이번 아이템은 __"기존 객체를 재사용해야 한다면 새로운 객체를 만들지 말라."__ 로 정리할 수 있다.