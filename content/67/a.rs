use std::collections::HashSet;

fn main() {
   let mut t = HashSet::new();
   t.insert(10);
   t.insert(11);
   for n in t {
      println!("{}", n);
   }
}
