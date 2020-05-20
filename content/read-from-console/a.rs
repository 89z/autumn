use std::io;
fn main() {
   let mut s1 = String::new();
   match io::stdin().read_line(&mut s1) {
      Err(e) => {
         dbg!(e);
      },
      Ok(n1) => {
         dbg!(n1, s1);
      }
   }
}
