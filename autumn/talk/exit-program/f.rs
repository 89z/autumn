use std::{
   char,
   error::Error
};

fn main() -> Result<(), Box<dyn Error>> {
   let n = 10;
   let c = char::from_digit(n, 16).ok_or(format!("{}", n))?;
   println!("{}", c);
   Ok(())
}
