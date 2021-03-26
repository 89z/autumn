use std::num;

fn main() -> Result<(), num::ParseIntError> {
   let s = "100";

   // example 1
   let n: u8 = s.parse()?;
   println!("{}", n == 100);

   // example 2
   let n = s.parse::<u8>()?;
   println!("{}", n == 100);

   Ok(())
}
