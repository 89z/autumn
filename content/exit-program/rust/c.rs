fn main() -> Result<(), u8> {
   let n = "AB".parse::<u8>().map_err(|e| {
      println!("{}", e);
      1
   })?;
   println!("{}", n);
   Ok(())
}
