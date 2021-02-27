use std::{
   fs::File,
   io
};

fn main() -> io::Result<()> {
   let mut open = File::open("b.rs")?;
   let mut create = File::create("c.rs")?;
   io::copy(&mut open, &mut create)?;
   Ok(())
}
