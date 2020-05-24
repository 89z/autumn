use std::error::Error;
use std::io;
fn main() -> Result<(), Box<dyn Error>> {
   let mut s1 = String::new();
   io::stdin().read_line(&mut s1)?;
   println!("{}", s1);
   Ok(())
}
