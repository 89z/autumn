use std::collections::HashSet;

fn main() {
   // example 1
   let mut t1 = HashSet::new();
   t1.insert(10);
   // example 2
   let mut t2 = HashSet::new();
   t2.insert("May");
   // print
   println!("{:?} {:?}", t1, t2);
}
