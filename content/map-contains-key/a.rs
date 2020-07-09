use std::collections::HashMap;
fn main() {
   let mut m = HashMap::new();
   m.insert("Sunday", 10);
   let b = m.contains_key("Sunday");
   println!("{}", b);
}
