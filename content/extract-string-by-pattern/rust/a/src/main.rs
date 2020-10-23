use regex::Regex;

fn main() -> Result<(), u8> {
   let o_re = Regex::new("a(..)").or_else(|e| {
      println!("{}", e);
      Err(1)
   })?;

   let o_cap = o_re.captures("January");
   println!("{:?}", o_cap);
   Ok(())
}
