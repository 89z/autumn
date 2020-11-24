use std::{
   fs::File,
   io,
   io::BufRead,
   io::BufReader
};

fn main() -> io::Result<()> {
   let e = File::open("a.rs")?;
   for e in BufReader::new(e).lines() {
      println!("{}", e?);
   }
   Ok(())
}
