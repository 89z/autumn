use std::time::{SystemTime, UNIX_EPOCH};
fn main() {
   let o = SystemTime::now();
   let r = o.duration_since(UNIX_EPOCH);
   println!("{:?}", r);
}
