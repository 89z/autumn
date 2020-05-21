use regex::Regex;
use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let s1 = "Wednesday";
   for a1 in Regex::new(r"e(..)")?.captures_iter(s1) {
      dbg!(a1);
   }
   Ok(())
}
