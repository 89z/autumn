use std::env;
fn main() {
   let mut a1 = env::args();
   let s1 = a1.nth(1);
   dbg!(s1);
}
