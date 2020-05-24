use regex::Regex;
use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let s1 = "Wednesday";
   let s2 = "e.";
   println!("example 1");
   let s3 = Regex::new(s2)?.find(s1).ok_or(s2)?.as_str();
   println!("{}", s3);
   println!("example 2");
   for o1 in Regex::new(s2)?.find_iter(s1) {
      let s3 = o1.as_str();
      println!("{}", s3);
   }
   Ok(())
}
