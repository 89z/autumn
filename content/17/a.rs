use std::{env, io};
fn main() -> io::Result<()> {
   let s = env::current_dir()?;
   println!("{:?}", s);
   Ok(())
}
