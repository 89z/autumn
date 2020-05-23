use std::env;
fn main() {
   let a1: Vec<String> = env::args().collect();
   println!("{:?}", a1);
}
