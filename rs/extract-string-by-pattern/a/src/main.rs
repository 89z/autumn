use regex::Regex;

fn main() -> Result<(), regex::Error> {
   let o_re = Regex::new("a(..)")?;
   let o_cap = o_re.captures("January");
   println!("{:?}", o_cap);
   Ok(())
}
