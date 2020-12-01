use std::{
   thread,
   time::Duration
};

fn main() {
   let o = Duration::from_secs(5);
   println!("sleep");
   thread::sleep(o);
}
