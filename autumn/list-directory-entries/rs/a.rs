use std::{fs, io};

fn main() -> io::Result<()> {
   let e_dir = fs::read_dir(".")?;
   for e_ent in e_dir {
      let o_ent = e_ent?;
      println!("{:?}", o_ent);
   }
   Ok(())
}
