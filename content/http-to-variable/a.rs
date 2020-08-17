use curl::easy::{Easy, WriteError};
use std::error::Error;
use std::io::{stdout, Write};
fn main() -> Result<(), Box<dyn Error>> {
   let mut o = Easy::new();
   o.url("https://speedtest.lax.hivelocity.net")?;
   o.write_function(|y|
      stdout().write(y).or(Err(WriteError::Pause))
   )?;
   o.perform()?;
   Ok(())
}
