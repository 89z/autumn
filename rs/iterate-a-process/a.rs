use std::{io, io::BufRead, io::Cursor, process::Command};

fn main() -> io::Result<()> {
   let o = Command::new("cc").arg("--print-search-dirs").output()?;

   for e in Cursor::new(o.stdout).lines() {
      println!("{}", e?);
   }

   Ok(())
}
