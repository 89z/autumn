use std::{
   fs,
   io
};

fn main() -> io::Result<()> {
   let m = fs::metadata("a.rs")?;
   let b = m.is_file();
   println!("{}", b);
   Ok(())
}
