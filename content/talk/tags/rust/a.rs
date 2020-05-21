use std::{env, process};
fn main() {
   let a = match env::var("BROWSER") {
      Err(b) => {
         dbg!(b);
         process::exit(1);
      }
      Ok(c) => c
   };
   dbg!(a);
}
