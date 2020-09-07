// example 1
use std::env;
// example 2
use std::env as o;
// end
fn main() {
   let a = env::args();
   let a2 = o::args();
   println!("{:?} {:?}", a, a2);
}
