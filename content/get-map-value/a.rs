use std::collections::HashMap;
fn main() {
   let mut m = HashMap::new();
   m.insert("Sunday", 10);
   let n = m["Sunday"];
   println!("{}", n);
}
