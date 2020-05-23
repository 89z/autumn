use std::env;
use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   // test 1
   let s1 = env::var("APPDATA")?;
   println!("{}", s1);
   // test 2
   let s2 = env::var("DATAAPP")?;
   println!("{}", s2);
   // test 3
   println!("Sunday");
   Ok(())
}
