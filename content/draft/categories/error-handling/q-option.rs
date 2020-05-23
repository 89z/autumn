use std::collections::HashMap;
use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let mut m = HashMap::new();
   m.insert("Sunday", 10);
   // std::error::Error is not implemented for std::option::NoneError
   let n = m.get("Sunday")?;
   println!("{}", n);
   Ok(())
}
