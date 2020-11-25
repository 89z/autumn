use std::{
   fs,
   io
};

fn main() -> io::Result<()> {
   let iter_e = fs::read_dir(".")?;
   for entry_e in iter_e {
      println!("{:?}", entry_e?);
   }
   Ok(())
}
