use std::time::{
   Duration,
   SystemTime,
   UNIX_EPOCH
};

fn unix(n: u64) -> Option<SystemTime> {
   let d = Duration::new(n, 0);
   UNIX_EPOCH.checked_add(d)
}

fn main() {
   let t = unix(0x5555_5555);
   println!("{:?}", t);
}
