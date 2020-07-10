use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let s = "10";
   let n: u8 = s.parse()?;
   println!("{}", n == 10);
   Ok(())
}
