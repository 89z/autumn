use std::{
   fs,
   io
};

fn main() -> io::Result<()> {
   let meta = fs::metadata("a.rs")?;
   let t = meta.modified()?;
   println!("{:?}", t);
   Ok(())
}
