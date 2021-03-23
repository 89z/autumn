use std::{
   fs,
   io
};

fn main() -> io::Result<()> {
   let iter = fs::read_dir(".")?;
   for entry in iter {
      println!("{:?}", entry?);
   }
   Ok(())
}
