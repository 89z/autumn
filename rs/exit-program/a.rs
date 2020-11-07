use std::num;

fn main() -> Result<(), num::ParseIntError> {
   let n: u8 = "AB".parse()?;
   println!("{}", n);
   Ok(())
}
