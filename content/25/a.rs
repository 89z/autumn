use std::num;

fn main() -> Result<(), num::ParseFloatError> {
   let s = "2.9";
   let n: f32 = s.parse()?;
   println!("{}", n == 2.9);
   Ok(())
}
