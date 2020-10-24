use std::error::Error;

fn main() -> Result<(), Box<dyn Error>> {
   let n: u8 = "AB".parse()?;
   println!("{}", n);
   Ok(())
}
