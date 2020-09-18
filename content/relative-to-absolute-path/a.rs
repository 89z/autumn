use std::{fs, io};

fn main() -> io::Result<()> {
   let sa = "index.md";
   let sb = fs::canonicalize(sa)?;
   println!("{:?}", sb);
   Ok(())
}
