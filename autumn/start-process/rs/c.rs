use std::{
   io,
   process::Command
};

fn main() -> io::Result<()> {
   Command::new("dust").status()?;
   Ok(())
}
