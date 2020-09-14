use std::fs::File;
use std::io::{BufRead, BufReader, Error};

fn main() -> Result<(), Error> {
   let r_file = File::open("a.txt")?;
   for r_line in BufReader::new(r_file).lines() {
      let s_line = r_line?;
      println!("{}", s_line);
   }
   Ok(())
}
