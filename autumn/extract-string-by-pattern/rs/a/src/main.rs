use regex::Regex;

fn main() -> Result<(), regex::Error> {
   let a = Regex::new("(..)n")?;
   if let Some(a) = a.captures("Sunday Monday") {
      println!("{}", &a[1] == "Su");
   }
   Ok(())
}
