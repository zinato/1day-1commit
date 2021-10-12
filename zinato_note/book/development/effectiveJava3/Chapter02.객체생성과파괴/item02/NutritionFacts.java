package item02;

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
