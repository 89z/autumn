use std::{env, io};
fn main() -> Result<(), io::Error> {
   let s = env::current_dir()?;
   println!("{:?}", s);
   Ok(())
}
