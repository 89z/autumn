use std::{fs, io};

fn main() -> io::Result<()> {
   for r_ent in fs::read_dir(".")? {
      let s_ent = r_ent?.path();
      println!("{:?}", s_ent);
   }
   Ok(())
}
