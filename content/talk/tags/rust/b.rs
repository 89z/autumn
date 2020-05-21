use std::env;
use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let a = env::var("SUNDAY")?;
   dbg!(a);
   Ok(())
}
