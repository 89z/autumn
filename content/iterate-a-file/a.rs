use std::fs::File;
use std::io::{BufRead,BufReader};
fn main() {
   match File::open("a.txt") {
      Err(e) => println!("{}", e),
      Ok(o1) => for o2 in BufReader::new(o1).lines() {
         match o2 {
            Err(e) => println!("{}", e),
            Ok(s1) => println!("{}", s1)
         }
      }
   }
}
