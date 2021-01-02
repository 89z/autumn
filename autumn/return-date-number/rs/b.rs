use std::time::{
   SystemTime,
   SystemTimeError,
   UNIX_EPOCH
};

fn main() -> Result<(), SystemTimeError> {
   let o = SystemTime::now().duration_since(UNIX_EPOCH)?;
   let n = o.as_secs();
   println!("{}", n);
   Ok(())
}
