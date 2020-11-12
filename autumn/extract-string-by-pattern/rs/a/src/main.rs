use regex::Regex;

fn main() -> Result<(), regex::Error> {
   let r = Regex::new("a(..)")?;
   let o = r.captures("January");
   println!("{:?}", o);
   Ok(())
}
