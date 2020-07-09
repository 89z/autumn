use std::{thread, time};
fn main() {
   println!("Sunday");
   thread::sleep(time::Duration::from_secs(5));
   println!("Monday");
}
