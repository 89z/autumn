use std::{io, process::Command};

fn main() -> io::Result<()> {
   let o = Command::new("cc").arg("--print-search-dirs").output()?;

   for s in String::from_utf8_lossy(&o.stdout).split_whitespace() {
      println!("{}", s);
   }

   Ok(())
}
