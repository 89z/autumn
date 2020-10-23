use regex::Regex;

fn main() -> Result<(), u8> {
   let o_re = Regex::new("a(..)").or_else(|e| {
      println!("{}", e);
      return Err(1);
   })?;

   for o_cap in o_re.captures_iter("January") {
      println!("{:?}", o_cap);
   }

   Ok(())
}
