use std::collections::HashMap;
fn main() {
   let mut m = HashMap::new();
   m.insert("Sunday", 10);
   if let Some(n) = m.get("Sunday") {
      println!("{}", n);
   }
}
