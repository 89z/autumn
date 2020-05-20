use std::collections::HashMap;
fn main() {
   let mut m1 = HashMap::new();
   m1.insert("Sun", 10);
   let n1 = m1.get("Sun");
   dbg!(n1);
}
