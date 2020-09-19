// example 1
use std::env;
// example 2
use std::env as o2;
// end
fn main() {
   let a = env::args();
   let a2 = o2::args();
   println!("{:?} {:?}", a, a2);
}
