use regex::Regex;

fn main() -> Result<(), String> {
   let o = Regex::new("(..)n").map_err(|e|
      e.to_string()
   )?;
   let a = o.captures("Sunday Monday").ok_or("captures")?;
   println!("{}", &a[1]);
   Ok(())
}
