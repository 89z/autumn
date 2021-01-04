use std::time::{
   Duration,
   UNIX_EPOCH
};

fn main() {
   let d = Duration::new(1609376551, 0);
   let t = UNIX_EPOCH.checked_add(d);
   println!("{:?}", t);
}
