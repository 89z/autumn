use std::borrow::Cow;

fn main() {
   let s1 = Cow::from("March");
   let s2 = Cow::from("April");
   println!("{}", s1 + s2);
}
