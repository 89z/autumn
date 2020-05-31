use std::env;
use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let o = env::current_dir()?;
   let s = o.display();
   println!("{}", s);
   Ok(())
}
