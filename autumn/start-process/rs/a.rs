use std::{
   io,
   process::Command
};

fn main() -> io::Result<()> {
   let mut c = Command::new(r"C:\Windows\notepad").spawn()?;
   c.wait()?;
   Ok(())
}
