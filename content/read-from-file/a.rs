use std::fs;
fn main() {
   match fs::read_to_string("a.txt") {
      Err(e) => println!("{}", e),
      Ok(s1) => println!("{}", s1)
   }
}
