fn main() -> Result<(), String> {
   Err(1).map_err(|e|
      e.to_string()
   )?;
   println!("March");
   Ok(())
}
