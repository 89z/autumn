use regex::Regex;
use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let s1 = "Wednesday";
   let mut s2;
   let mut s3;
   // test 1
   s2 = "e.";
   s3 = Regex::new(s2)?.find(s1).ok_or(s2)?.as_str();
   println!("{}", s3);
   // test 2
   s2 = "f.";
   s3 = Regex::new(s2)?.find(s1).ok_or(s2)?.as_str();
   println!("{}", s3);
   // test 3
   println!("Ok");
   Ok(())
}
