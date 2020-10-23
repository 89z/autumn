use curl::easy::{Easy, WriteError};
use std::io::{stdout, Write};

fn main() -> Result<(), u8> {
   let mut o = Easy::new();

   o.url("https://speedtest.lax.hivelocity.net").or_else(|e| {
      println!("{}", e);
      Err(1)
   })?;

   o.write_function(|y| {
      stdout().write(y).or(Err(WriteError::Pause))
   }).or_else(|e| {
      println!("{}", e);
      Err(1)
   })?;

   o.perform().or_else(|e| {
      println!("{}", e);
      Err(1)
   })?;

   Ok(())
}
