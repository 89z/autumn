use std::{fs, io};

fn main() -> io::Result<()> {
   let r = fs::read_dir(".")?;
   for r in r {
      println!("{:?}", r?);
   }
   Ok(())
}
