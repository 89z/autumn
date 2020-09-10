use std::{fs, io};

fn main() -> io::Result<()> {
   let s= "index.md";
   let s2 = fs::canonicalize(s)?;
   println!("{:?}", s2);
   Ok(())
}
