use std::collections::HashSet;

fn main() {
   let mut t = HashSet::new();
   t.insert("May");
   t.insert("June");
   for s in t {
      println!("{}", s);
   }
}
