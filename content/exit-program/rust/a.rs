use std::error::Error;

fn main() -> Result<(), Box<dyn Error>> {
   let n = "AB".parse::<u8>()?;
   println!("{}", n);
   Ok(())
}
