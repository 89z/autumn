use regex::Regex;
use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let s1 = "Wednesday";
   let s2 = "e(..)";
   println!("example 1");
   let o1 = Regex::new(s2)?.captures(s1).ok_or(s2)?;
   println!("{:?}", o1);
   println!("example 2");
   for o1 in Regex::new(s2)?.captures_iter(s1) {
      println!("{:?}", o1);
   }
   Ok(())
}
