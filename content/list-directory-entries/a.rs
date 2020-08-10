use std::error::Error;
use std::fs;
fn main() -> Result<(), Box<dyn Error>> {
   for o_read in fs::read_dir(".")? {
      let o_path = o_read?.path();
      let s = o_path.display();
      println!("{}", s);
   }
   Ok(())
}
