use std::env;
fn main() {
   let mut o1 = env::args();
   let o2 = o1.nth(1);
   println!("{:?}", o2);
}
