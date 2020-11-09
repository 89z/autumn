use std::borrow::Cow;

fn main() {
   let s = Cow::Borrowed("March");
   println!("{}", s);
}
