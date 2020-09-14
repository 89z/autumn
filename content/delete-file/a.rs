use std::{fs, io};

fn main() -> io::Result<()> {
   fs::remove_file("a.txt")?;
   Ok(())
}
