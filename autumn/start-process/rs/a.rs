use std::{io, process::Command};

fn main() -> io::Result<()> {
   println!("BEGIN");
   let mut o = Command::new("pipe").spawn()?;
   o.wait()?;
   Ok(())
}
