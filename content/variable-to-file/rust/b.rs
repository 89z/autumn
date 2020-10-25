use std::{fs::File, io::Error, io::Write};

fn main() -> Result<(), Error> {
   let y = b"May\n";
   let mut o = File::create("a.txt")?;
   o.write(y)?;
   Ok(())
}
