# ITEM 02 : 생성자에 매개변수가 많다면 빌더를 고려하라

- 정적 팩토리 메서드와 생성자에는 선택적 매개변수가 많을 때 적절히 대응하기 어렵다는 단점이 있다.
- 이럴 때 점층적 생성자 패턴(telescoping constructor patter)을 많이 사용했다.

## 점층적 생성자 패턴 (Telescoping Contructor Pattern)

```java
// 점층적 생성자 패턴 - 확장하기 어렵다. 
public class NutritionFacts {
  private final int servingSize;  
  private final int servings;     
  private final int calories;     
  private final int fat;
  private final int soldium;
  private final int carbohydrate;

  public NutritionFacts(int servingSize, int servings) {
    this(servingSize, servings, 0);
  }

  public NutritionFacts(int servingSize, int servings, int calories) {
    this(servingSize, servings, calories, 0);
  }

  public NutritionFacts(int servingSize, int servings, int calories, int fat) {
    this(servingSize, servings, calories, fat, 0);
  }

  public NutritionFacts(int servingSize, int servings, int calories, int fat, int soldium,
  int carbohydrate) {
    this.servingSize = servingSize;
    this.servings = servings;
    this.calories = calories;
    this.fat = fat;
    this.soldium = soldium;
    this.carbohydrate = carbohydrate;
  }

  public static void main(String[] args) {
    NutritionFacts cocaCola = new NutritionFacts(240, 8, 100, 0, 35, 27);
  }
  
}
```

- 설정하지 원하지 않는 매개변수도 값을 설정해줘야 한다.
- 매개변수 개수가 늘어나면 클라이언트 코드를 작성하거나 읽기 어려워 진다. 실수할 확률이 올라간다.
  - 순서를 바꿔서 입력해도 컴파일러에서는 어떠한 에러도 일어나지 않는다.
  - 값이 무슨 의미인지 알기가 어렵다.
  - 개수도 주의 깊게 세어봐야 한다.
- 자바 빈즈 패턴을 통해 해결할 수 있다.
  - 매개변수가 없는 생성자로 객체를 만든 후, 세터(setter) 메서드들을 호출해 원하는 매개변수의 값을 설정하는 방식이다.

## 자바 빈즈 패턴 (Java Beans Pattern)

```java
public class NutrionsFactsJavaBean {

  //매개 변수들은 (기본값이 있다면) 기본값을 초기화 한다.
  private int servingSize = -1;  
  private int servings = -1;     
  private int calories = 0;     
  private int fat = 0;
  private int soldium = 0;
  private int carbohydrate = 0;

  public NutrionsFactsJavaBean() {}
  //setter method
  public void setServingSize(int val) { servingSize = val; }
  public void setServings(int val) { servings = val; }
  public void setCalories(int val) { calories = val; }
  public void setFat(int val) { fat = val; }
  public void setSoldium(int val) {soldium = val; }
  public void setCarbohydrate(int val) {carbohydrate = val; }

  public static void main(String[] args) {
    NutrionsFactsJavaBean cocaCola = new NutrionsFactsJavaBean();
    cocaCola.setServingSize(240);
    cocaCola.setServings(8);
    cocaCola.setCalories(100);
    cocaCola.setSoldium(35);
    cocaCola.setCarbohydrate(27);
  }
  
}
```

- 자바 빈즈 패턴은 자신만의 심각한 단점을 가지고 있다.
  1. 객체 하나를 만들려면 메서드를 여러개 호출해야 한다.
  2. 객체가 완전히 생성되기 전까지는 일관성이 무너진 상태이다.
      - 점층적 생성자 패턴에서는 매개변수들이 유효한지 생성자에서만 확인하면 일관성을 유지할 수 있었지만 자바빈즈는 확인할 수 없다.

> 점층적 생성자 패턴의 안정성과 자바 빈즈 패턴의 가독성의 장점을 갖춘 것이 ***빌더 패턴(Builder pattern)*** 이다.

## 빌더 패턴 (Builder pattern)
- [코드 참고](https://github.com/zinato/1day-1commit/blob/zinato/2021_08/zinato_note/book/development/EffectiveJava3/item02/NutritionFactsBuilderPattern.java)

```java
//파이썬과 스칼라에 있는 네임드 매개변수를 사용하여 가독성을 높일 수 있다.
    NutriotionFactsBuilderPattern cocaCola = new NutriotionFactsBuilderPattern.Builder(240, 8) //필수 매개변수 
      .calories(100) //메서드 체인닝 
      .sodium(35)
      .carbohydrate(27)
      .build();

```

- 빌더의 세터 메서드들을 빌더 자신을 반환하기 때문에 연쇄적으로 호출할 수 있다. (method chaining)
- 빌더 패턴은 파이썬과 스칼라에 있는 named optional parameter를 흉내낸 것이다.
- 점층적 생성자 패턴보다 코드가 깔끔하고 가독성이 있으며 필수 매개변수와 옵션 매개변수를 구분하여 저장할 수 있다.
- 클라이언트가 코드를 읽기 쉬고 쓰기 쉽다.
- 서버 쪽 코딩이 조금 더 투자된다.

## 계층적으로 설계된 클래스와 잘 어울리는 빌더 패턴
