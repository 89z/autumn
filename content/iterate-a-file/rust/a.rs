use std::{fs::File, io::BufRead, io::BufReader, io::Error};

fn main() -> Result<(), Error> {
   let e_file = File::open("index.md")?;
   for e_line in BufReader::new(e_file).lines() {
      println!("{}", e_line?);
   }
   Ok(())
}
