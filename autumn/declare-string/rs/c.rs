use std::borrow::Cow;

fn main() {
   let s = Cow::from("March");
   println!("{}", s);
}
