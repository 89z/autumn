use std::{
   io::BufRead,
   io::Cursor,
   io::Error,
   process::Command
};

fn main() -> Result<(), Error> {
   let o = Command::new("cc").arg("--print-search-dirs").output()?;
   for e in Cursor::new(o.stdout).lines() {
      println!("{}", e?);
   }
   Ok(())
}
