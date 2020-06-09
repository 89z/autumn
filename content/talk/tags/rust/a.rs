use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let o = Some(10);
   let n = o.ok_or("Err")?;
   println!("{}", n);
   Ok(())
}
