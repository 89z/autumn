use std::num;

fn main() -> Result<(), num::ParseIntError> {
   let s = "q3ezbz";
   let n = u32::from_str_radix(s, 36)?;
   println!("{}", n == 1609480799);
   Ok(())
}
