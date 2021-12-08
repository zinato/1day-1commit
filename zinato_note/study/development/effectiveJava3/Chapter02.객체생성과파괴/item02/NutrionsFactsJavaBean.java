package item02;

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
