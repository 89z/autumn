use std::{
   fs::File,
   io
};

fn main() -> io::Result<()> {
   let mut in_o = File::open("b.rs")?;
   let mut out_o = File::create("c.rs")?;
   io::copy(&mut in_o, &mut out_o)?;
   Ok(())
}
