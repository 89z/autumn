use curl::easy::{Easy, WriteError};
use std::io::{stdout, Write};
fn main() -> Result<(), curl::Error> {
   let mut o = Easy::new();
   o.url("https://speedtest.lax.hivelocity.net")?;
   o.write_function(|y| {
      if let Ok(n) = stdout().write(y) {
         return Ok(n)
      }
      Err(WriteError::Pause)
   })?;
   o.perform()?;
   Ok(())
}
