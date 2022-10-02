fn main() {
    let fruit = vec!["apple", "banana", "grape"];
    let buffer_overflow = fruit[4];
    assert_eq!(buffer_overflow, "watermelon");
}
