use std::collections::HashSet;

fn main() {
   let mut t = HashSet::new();
   t.insert(10);
   let b = t.contains(&10);
   println!("{}", b);
}
