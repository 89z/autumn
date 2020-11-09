use std::borrow::Cow;

fn main() {
   let s1 = Cow::Borrowed("March");
   let s2 = Cow::Borrowed("April");
   println!("{}", s1 + s2);
}
