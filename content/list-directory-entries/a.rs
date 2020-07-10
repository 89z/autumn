use std::error::Error;
use std::fs;
fn main() -> Result<(), Box<dyn Error>> {
   for o1 in fs::read_dir(".")? {
      let o2 = o1?.path();
      let s1 = o2.display();
      println!("{}", s1);
   }
   Ok(())
}
