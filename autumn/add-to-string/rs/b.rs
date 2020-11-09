use std::borrow::Cow;

fn main() {
   let mut s = Cow::Borrowed("March");
   s += "April";
   println!("{}", s);
}
