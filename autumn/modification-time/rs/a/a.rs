use std::{
   fs,
   io
};

fn main() -> io::Result<()> {
   let e = fs::metadata("a.rs")?;
   let e = e.modified()?;
   println!("{:?}", e);
   Ok(())
}
