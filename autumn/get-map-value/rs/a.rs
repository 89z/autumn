use std::collections::HashMap;

fn main() {
   let mut m = HashMap::new();
   m.insert("day", 31);
   let n = m["day"];
   println!("{}", n == 31);
}
