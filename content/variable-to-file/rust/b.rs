use std::{fs::File, io, io::Write};

fn main() -> io::Result<()> {
   let y = b"March\n";
   let mut o = File::create("a.txt")?;
   o.write(y)?;
   Ok(())
}
