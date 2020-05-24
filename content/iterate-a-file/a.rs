use std::error::Error;
use std::fs::File;
use std::io::{BufRead, BufReader};
fn main() -> Result<(), Box<dyn Error>> {
   let o1 = File::open("a.txt")?;
   for o2 in BufReader::new(o1).lines() {
      let s1 = o2?;
      println!("{}", s1);
   }
   Ok(())
}
