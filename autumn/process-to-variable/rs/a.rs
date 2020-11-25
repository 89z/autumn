use std::{
   io,
   io::Write,
   process::Command
};

fn main() -> io::Result<()> {
   let o = Command::new("rustc").arg("-V").output()?;
   io::stdout().write(&o.stdout)?;
   Ok(())
}
