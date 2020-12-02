use std::{
   fs::File,
   io::BufRead,
   io::BufReader
};

fn main() -> std::io::Result<()> {
   let e = File::open("a.rs")?;
   for e in BufReader::new(e).lines() {
      println!("{}", e?);
   }
   Ok(())
}
