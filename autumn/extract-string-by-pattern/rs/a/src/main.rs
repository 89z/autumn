use {
   regex::Regex,
   std::error::Error
};

fn main() -> Result<(), Box<dyn Error>> {
   let o = Regex::new("(..)n")?;
   let s = "Sunday Monday";
   let a = o.captures(s).ok_or(s)?;
   println!("{}", &a[1]);
   Ok(())
}
