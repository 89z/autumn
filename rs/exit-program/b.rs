fn main() -> Result<(), String> {
   let n = "AB".parse::<u8>().map_err(|e|
      e.to_string()
   )?;
   println!("{}", n);
   Ok(())
}
