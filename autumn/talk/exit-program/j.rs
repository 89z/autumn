use std::{
   char,
   error::Error
};

fn main() -> Result<(), Box<dyn Error>> {
   let n = 10;
   let c = match char::from_digit(n, 16) {
      Some(v) => v,
      None => Err(format!("{}", n))?
   };
   println!("{}", c);
   Ok(())
}
