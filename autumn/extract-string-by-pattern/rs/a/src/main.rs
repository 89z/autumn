use {
   regex::Error,
   regex::Regex
};

fn main() -> Result<(), Error> {
   let r = Regex::new("(..)n")?;
   let c = r.captures("Sunday Monday");  
   println!("{:?}", c);
   Ok(())
}
