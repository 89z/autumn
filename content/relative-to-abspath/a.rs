use std::{fs, io};
fn main() -> io::Result<()> {
   let s1 = "index.md";
   let s2 = fs::canonicalize(s1)?;
   println!("{:?}", s2);
   Ok(())
}
