package item02.pizza;

public abstract class Pizza {
  public enum Topping {
    HAM, MUSHROOM, ONION, PEPPER, SAUSAGE
  }
  final Set<Topping> toppings;

  // 재귀적 타입 한정을 이용하는 제네릭 타입
  // 추상 메소드 self를 사용하여 메서드 체이닝을 지원 (simulated self-type)

  abstract static class Builder<T extends Builder<T>> {
    EnumSet<Topping> toppings = EnumSet.noneOf(Topping.class);
    public T addTopping(Topping topping) {
      toppings.add(Objects.requireNonNull(topping));
      return self();
    }

    abstract Pizza build();

    // 하위 클래스는 이 메서드를 오버라이드하여
    // "this"를 반환하도록 해야 한다.
    protected abstract T self();
  }

  Pizza(Builder<?> builder) {
    toppings = builder.toppings.clone();
  }
}
