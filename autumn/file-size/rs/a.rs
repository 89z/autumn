use std::{
   fs,
   io
};

fn main() -> io::Result<()> {
   let o = fs::metadata("a.rs")?;
   let n = o.len();
   println!("{}", n);
   Ok(())
}
