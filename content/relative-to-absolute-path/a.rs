use std::{fs, io};

fn main() -> io::Result<()> {
   let s = "index.md";
   let s1 = fs::canonicalize(s)?;
   println!("{:?}", s1);
   Ok(())
}
