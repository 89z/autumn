use std::num;
fn main() -> Result<(), num::ParseFloatError> {
   let s = "1.9";
   let n: f32 = s.parse()?;
   println!("{}", n == 1.9);
   Ok(())
}
