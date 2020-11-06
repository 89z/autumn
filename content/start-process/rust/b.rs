use std::{io, process::Command};

fn main() -> io::Result<()> {
   Command::new("waterfox").arg("google.com/search?q=march").spawn()?;
   Ok(())
}
