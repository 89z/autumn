use std::process;

fn main() {
   let n: u8 = match Err(1) {
      Ok(v) => v,
      Err(e) => {
         println!("{}", e);
         process::exit(1);
      }
   };
   println!("{}", n);
}
