use regex::Regex;

fn main() -> Result<(), regex::Error> {
   let o_re = Regex::new("a(..)")?;
   for o_cap in o_re.captures_iter("January") {
      println!("{:?}", o_cap);
   }
   Ok(())
}
