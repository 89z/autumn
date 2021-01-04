use std::collections::HashMap;

fn main() {
   let mut m = HashMap::new();
   m.insert("day", 4);
   let b = m.contains_key("day");
   println!("{}", b);
}
