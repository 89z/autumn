use std::collections::HashMap;
use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let mut m = HashMap::new();
   m.insert("Sunday", 10);
   let n = m.get("Sunday").ok_or("None")?;
   println!("{}", n);
   Ok(())
}
