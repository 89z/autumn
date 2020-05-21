use std::env;
fn main() {
   if let Ok(s1) = env::var("BROWSER") {
      dbg!(s1);
   }
}
