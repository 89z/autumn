use std::{fs, io};

fn main() -> io::Result<()> {
   fs::copy("a.txt", "b.txt")?;
   Ok(())
}
