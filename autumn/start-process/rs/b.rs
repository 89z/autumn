use std::{
   io,
   process::Command
};

fn main() -> io::Result<()> {
   let c = Command::new("dust").arg("-V").output()?;
   println!("{}", c.stdout == b"Dust 0.5.4\n");
   Ok(())
}
