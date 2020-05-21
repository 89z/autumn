use std::io;
fn main() {
   let mut s1 = String::new();
   if let Ok(n1) = io::stdin().read_line(&mut s1) {
      dbg!(n1, s1);
   }
}
