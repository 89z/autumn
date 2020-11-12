use std::{fs, io};

fn main() -> io::Result<()> {
   let r = fs::metadata("a.rs")?;
   let r = r.modified()?;
   println!("{:?}", r);
   Ok(())
}
