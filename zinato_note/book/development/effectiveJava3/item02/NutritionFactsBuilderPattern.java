package item02;

public class NutritionFactsBuilderPattern {
  //2-3 빌더 패턴 - 점층적 생성자 패턴과 자바빈즈 패턴의 장점만 취했다.
  private final int servingSize;  
  private final int servings;     
  private final int calories;     
  private final int fat;
  private final int sodium;
  private final int carbohydrate;

  public static class Builder {
    //필수 매개변수
    private final int servingSize;
    private final int servings;

    //선택 매개변수 - 기본값으로 초기화한다.
    private int calories = 0;
    private int fat = 0;
    private int sodium = 0;
    private int carbohydrate = 0;

    public Builder(int servingSize, int servings) {
      this.servingSize = servingSize;
      this.servings = servings;
    }

    public Builder calories(int val) {
      calories = val;
      return this;
    }

    public Builder fat(int val) {
      fat = val;
      return this;
    }

    public Builder sodium(int val) {
      sodium = val;
      return this;
    }

    public Builder carbohydrate(int val) {
      carbohydrate = val;
      return this;
    }

    public NutritionFactsBuilderPattern build() {
      return new NutritionFactsBuilderPattern(this);
    }
  }

  private NutritionFactsBuilderPattern(Builder builder) {
    servingSize = builder.servingSize;
    servings = builder.servings;
    calories = builder.calories;
    fat = builder.fat;
    sodium = builder.sodium;
    carbohydrate = builder.carbohydrate;
  }

  public static void main(String[] args) {
    //파이썬과 스칼라에 있는 네임드 매개변수를 사용하여 가독성을 높일 수 있다.
    NutritionFactsBuilderPattern cocaCola = new NutritionFactsBuilderPattern.Builder(240, 8) //필수 매개변수 
      .calories(100) //메서드 체인닝 
      .sodium(35)
      .carbohydrate(27)
      .build();
  }

}
