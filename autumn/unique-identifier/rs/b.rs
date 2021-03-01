use std::num;

fn main() -> Result<(), num::ParseIntError> {
   let s = "1z141z3";
   let n = u32::from_str_radix(s, 36)?;
   println!("{}", n == 0xFFFF_FFFF);
   Ok(())
}
