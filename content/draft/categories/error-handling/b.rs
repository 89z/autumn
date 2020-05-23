use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let s1 = "Sunday Monday";
   let mut a1 = s1.split_whitespace();
   // std::error::Error is not implemented for std::option::NoneError
   let s2 = a1.nth(1)?;
   println!("{}", s2);
   Ok(())
}
