package item02.pizza;

public class NyPizza extends Pizza{
  public enum Size {
    SMALL, MEDIUM, LARGE
  }
  private final Size size;

  public static class Bulder extends Pizza.Builder<Builder> {
    private final Size size;

    public Builder(Size size) {
      this.size = Objects.requireNonNull(size);
    }

    @Override
    public NyPizze builder() {
      return new NyPizza(this);
    }

    @Override
    protected Builder self() {
      return this;
    }

    private NyPizza(Builder builder) {
      super(builder);
      size = builder.size;
    }
  }
}
