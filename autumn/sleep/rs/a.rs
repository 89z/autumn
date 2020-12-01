use std::{
   thread,
   time::Duration
};

fn main() {
   let o = Duration::from_millis(500);
   loop {
      println!("sleep");
      thread::sleep(o);
   }
}
