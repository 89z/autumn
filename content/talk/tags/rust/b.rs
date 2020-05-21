use std::{env, error};
fn main() -> Result<(), Box<dyn error::Error>> {
   let a = env::var("SUNDAY")?;
   dbg!(a);
   Ok(())
}
