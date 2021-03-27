use std::{
   io::BufRead,
   io::Cursor,
   process::Command
};

fn main() -> std::io::Result<()> {
   let cmd = Command::new("cc").arg("--print-search-dirs").output()?;
   for each in Cursor::new(cmd.stdout).lines() {
      println!("{}", each?);
   }
   Ok(())
}
