use std::borrow::Cow;

fn main() {
   let s = Cow::from("north");
   println!("{}", s);
}
