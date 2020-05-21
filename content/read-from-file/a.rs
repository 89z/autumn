use std::fs;
fn main() {
   if let Ok(s1) = fs::read_to_string("a.txt") {
      dbg!(s1);
   }
}
