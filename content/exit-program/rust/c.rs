use std::process;

fn main() {
   let n = match "AB".parse::<u8>() {
      Ok(v) => v,
      Err(e) => {
         println!("{}", e);
         process::exit(1);
      }
   };
   println!("{}", n);
}
