use std::num;
fn main() -> Result<(), num::ParseIntError> {
   let s = "10";
   let n: u8 = s.parse()?;
   println!("{}", n == 10);
   Ok(())
}
