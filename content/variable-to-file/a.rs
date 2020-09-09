use std::{io, fs};

fn main() -> io::Result<()> {
   let s = "May\n";
   fs::write("a.txt", s)?;
   Ok(())
}
