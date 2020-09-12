use std::collections::HashMap;

fn main() {
   let mut m = HashMap::new();
   m.insert("year", 2019);
   let n = m["year"];
   println!("{}", n == 2019);
}
