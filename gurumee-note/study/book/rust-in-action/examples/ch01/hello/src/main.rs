fn greet_world() {
    let _english = "Hello World!";
    let _korean = "안녕 세상아!";
    let _japanese = "こんにちは世界";
    let regions = [_english, _korean, _japanese];

    for region in regions.iter() {
        println!("{}", &region);
    }
}

fn main() {
    greet_world();
}
