use std::{
   char,
   error::Error,
   fs
};

fn main() -> Result<(), Box<dyn Error>> {
   // example 1
   let o = fs::metadata("a.rs")?;
   // example 2
   let n = 10;
   let c = char::from_digit(n, 16).ok_or(format!("{}", n))?;
   // print
   println!("{} {}", o.len(), c);
   Ok(())
}
