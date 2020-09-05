use std::fs::File;
use std::io::{Error, Write};
fn main() -> Result<(), Error> {
   let mut o = File::create("a.txt")?;
   // example 1
   o.write(b"Sunday\n")?;
   // example 2
   write!(o, "Sunday\n")?;
   // return
   Ok(())
}
