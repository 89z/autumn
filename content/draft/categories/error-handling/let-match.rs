use regex::Regex;
use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let s1 = "Wednesday";
   let mut s2;
   let mut s3;
   // test 1
   s2 = "e.";
   s3 = match Regex::new(s2)?.find(s1) {
      Some(o) => o.as_str(),
      None => return Err(Box::from(s2))
   };
   println!("{}", s3);
   // test 2
   s2 = "f.";
   s3 = match Regex::new(s2)?.find(s1) {
      Some(o) => o.as_str(),
      None => return Err(Box::from(s2))
   };
   println!("{}", s3);
   // test 3
   println!("Ok");
   Ok(())
}
