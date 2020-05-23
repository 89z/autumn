use std::collections::HashMap;
fn main() {
   let mut m = HashMap::new();
   m.insert("Sunday", 10);
   match m.get("Sunday") {
      None => return,
      Some(n) => println!("{}", n)
   }
}
