use {
   regex::Error,
   regex::Regex
};

fn main() -> Result<(), Error> {
   let r = Regex::new("o(..)")?;
   let c = r.captures("south north");
   println!("{:?}", c);
   Ok(())
}
