use std::process;

fn main() {
   let n: u8 = match "AB".parse() {
      Ok(v) => v,
      Err(e) => {
         println!("{}", e);
         process::exit(1);
      }
   };
   println!("{}", n);
}
