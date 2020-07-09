use std::error::Error;
use std::fs;
fn main() -> Result<(), Box<dyn Error>> {
   fs::write("a.txt", "Sunday\n")?;
   Ok(())
}
