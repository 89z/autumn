use std::{io, fs};
fn main() -> io::Result<()> {
   fs::write("a.txt", "Sunday\n")?;
   Ok(())
}
