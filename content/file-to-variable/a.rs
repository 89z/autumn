use std::error::Error;
use std::fs;
fn main() -> Result<(), Box<dyn Error>> {
   let s1 = fs::read_to_string("a.txt")?;
   println!("{}", s1);
   Ok(())
}
