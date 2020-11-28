fn main() -> Result<(), u8> {
   Err(false).map_err(|e| {
      println!("{}", e);
      1
   })?;
   println!("March");
   Ok(())
}
