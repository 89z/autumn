use regex::Regex;

fn main() -> Result<(), regex::Error> {
   let r = Regex::new("o(..)")?;
   for c in r.captures_iter("south north") {
      println!("{:?}", c);
   }
   Ok(())
}
