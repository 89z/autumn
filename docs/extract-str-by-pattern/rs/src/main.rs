use regex::Regex;
use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let s1 = "Wednesday";
   // example 1
   match Regex::new(r"e.")?.find(s1) {
      None => println!("None"),
      Some(o1) => println!("{}", o1.as_str())
   }
   // example 2
   for o1 in Regex::new(r"e.")?.find_iter(s1) {
      println!("{}", o1.as_str());
   }
   // example 3
   match Regex::new(r"e(..)")?.captures(s1) {
      None => println!("None"),
      Some(o1) => println!("{:?}", o1)
   }
   // example 4
   for o1 in Regex::new(r"e(..)")?.captures_iter(s1) {
      println!("{:?}", o1);
   }
   Ok(())
}
