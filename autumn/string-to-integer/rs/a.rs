use std::num;

fn main() -> Result<(), num::ParseIntError> {
   let s = "100";
   // example 1
   let n1: u8 = s.parse()?;
   // example 2
   let n2 = s.parse::<u8>()?;
   // print
   println!("{}", n1 == 100 && n2 == 100);
   Ok(())
}
