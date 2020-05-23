use std::env;
use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let s1 = env::var("BROWSER")?;
   println!("{}", s1);
   Ok(())
}
