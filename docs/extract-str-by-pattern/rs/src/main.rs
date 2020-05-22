use regex::Regex;
use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let s1 = "Wednesday";
   // example 1
   // example 2
   for o1 in Regex::new(r"e.")?.find_iter(s1) {
      let s2 = o1.as_str();
      dbg!(s2);
   }
   // example 3
   // example 4
   for o1 in Regex::new(r"e(..)")?.captures_iter(s1) {
      dbg!(o1);
   }
   Ok(())
}
