use regex::{Error, Regex};

fn main() -> Result<(), Error> {
   let o_re = Regex::new("a(..)")?;
   // example 1
   let o_cap = o_re.captures("January");
   println!("{:?}", o_cap);
   // example 2
   for o_cap in o_re.captures_iter("January") {
      println!("{:?}", o_cap);
   }
   // return
   Ok(())
}
