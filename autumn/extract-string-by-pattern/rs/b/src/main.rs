use regex::Regex;

fn main() -> Result<(), regex::Error> {
   let r = Regex::new("a(..)")?;
   for o in r.captures_iter("January") {
      println!("{:?}", o);
   }
   Ok(())
}
