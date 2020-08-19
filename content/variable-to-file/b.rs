use std::error::Error;
use std::fs::File;
use std::io::Write;
fn main() -> Result<(), Box<dyn Error>> {
   let mut o1 = File::create("a.txt")?;
   // example 1
   o1.write(b"Sunday\n")?;
   // example 2
   write!(o1, "Sunday\n")?;
   // return
   Ok(())
}
