use std::{fs::File, io, io::BufRead, io::BufReader};

fn main() -> io::Result<()> {
   let r = File::open("a.rs")?;
   for r in BufReader::new(r).lines() {
      println!("{}", r?);
   }
   Ok(())
}
