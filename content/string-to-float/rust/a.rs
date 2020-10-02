use std::num;

fn main() -> Result<(), num::ParseFloatError> {
   let s = "9.9";
   let n: f32 = s.parse()?;
   println!("{}", n == 9.9);
   Ok(())
}
