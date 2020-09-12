use std::{fs, io};

fn main() -> io::Result<()> {
   let s = fs::read_to_string("a.txt")?;
   print!("{}", s);
   Ok(())
}
