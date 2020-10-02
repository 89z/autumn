use std::collections::HashSet;

fn main() {
   let mut t = HashSet::new();
   t.insert("May");
   t.insert("June");
   let b = t.contains("May");
   println!("{}", b);
}
