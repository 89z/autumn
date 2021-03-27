use regex::Regex;

fn main() -> Result<(), regex::Error> {
   let r = Regex::new("a(..)")?;
   for c in r.captures_iter("January") {
      println!("{:?}", c);
   }
   Ok(())
}
