use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let s = "1.9";
   let n: f32 = s.parse()?;
   println!("{}", n == 1.9);
   Ok(())
}
