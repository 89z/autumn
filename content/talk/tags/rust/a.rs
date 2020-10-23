fn main() -> Result<(), u8> {
   // example 1
   let n: u8 = "11".parse().or_else(|e| {
      println!("{}", e);
      Err(1)
   })?;
   println!("{}", n);
   // example 2
   let n: u8 = "AB".parse().or_else(|e| {
      println!("{}", e);
      Err(1)
   })?;
   println!("{}", n);
   // return
   Ok(())
}
