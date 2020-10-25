use std::{fs, io};

fn main() -> Result<(), io::Error> {
   let o_meta = fs::metadata("index.md")?;
   let o_mod = o_meta.modified()?;
   println!("{:?}", o_mod);
   Ok(())
}
