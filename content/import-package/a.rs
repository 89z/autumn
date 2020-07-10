// example 1
use std::env;
// example 2
use std::env as o1;
// end
fn main() {
   let a1 = env::args();
   let a2 = o1::args();
   dbg!(a1, a2);
}
