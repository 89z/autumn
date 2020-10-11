use std::{env, io};

fn main() -> io::Result<()> {
   env::set_current_dir("..")?;
   let o = env::current_dir()?;
   println!("{:?}", o);
   Ok(())
}
