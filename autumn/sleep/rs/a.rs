use std::{
   thread,
   time::Duration
};

fn main() {
   let d = Duration::from_millis(500);
   loop {
      println!("sleep");
      thread::sleep(d);
   }
}
