use std::time::{
   SystemTime,
   SystemTimeError,
   UNIX_EPOCH
};

fn main() -> Result<(), SystemTimeError> {
   let now_o = SystemTime::now();
   let dur_o = now_o.duration_since(UNIX_EPOCH)?;
   let n = dur_o.as_secs();
   println!("{}", n);
   Ok(())
}
