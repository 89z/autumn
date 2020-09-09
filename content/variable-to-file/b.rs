use std::fs::File;
use std::io::{Error, Write};

fn main() -> Result<(), Error> {
   let y = b"May\n";
   let mut o = File::create("a.txt")?;
   o.write(y)?;
   Ok(())
}
