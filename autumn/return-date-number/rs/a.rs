use std::time::{
   SystemTime,
   SystemTimeError,
   UNIX_EPOCH
};

fn main() -> Result<(), SystemTimeError> {
   let t = SystemTime::now().duration_since(UNIX_EPOCH)?;
   let n = t.as_millis();
   println!("{}", n);
   Ok(())
}
