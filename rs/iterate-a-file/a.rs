use std::{
   fs::File,
   io,
   io::BufRead,
   io::BufReader
};

fn main() -> io::Result<()> {
   let e_file = File::open("index.md")?;
   for e_line in BufReader::new(e_file).lines() {
      println!("{}", e_line?);
   }
   Ok(())
}
