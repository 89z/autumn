use std::{
   fs,
   io
};

fn main() -> io::Result<()> {
   let o = fs::metadata("a.rs")?;
   let b = o.is_file();
   println!("{}", b);
   Ok(())
}
