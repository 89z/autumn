use std::collections::HashMap;

fn main() {
   let mut m = HashMap::new();
   m.insert("year", 2019);
   let b = m.contains_key("year");
   println!("{}", b);
}
