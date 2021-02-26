use std::borrow::Cow;

fn main() {
   let mut s = Cow::from("March");
   s += "April";
   println!("{}", s);
}
