use pico_args::{Arguments, Error};

fn main() -> Result<(), Error> {
   let mut o = Arguments::from_env();
   let s_start = o.opt_value_from_str("--start")?.unwrap_or(String::new());
   let s_end = o.opt_value_from_str("--end")?.unwrap_or(String::new());
   let a = o.free()?;
   if a.len() != 1 {
      let a = Error::ArgumentParsingFailed{cause:String::from("free")};
      return Err(a);
   }
   Ok(())
}
