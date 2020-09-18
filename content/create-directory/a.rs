use std::fs;

fn main() {
   // example 1
   let e1 = fs::create_dir("May");
   // example 2
   let e2 = fs::create_dir_all("June/July");
   // print
   println!("{:?} {:?}", e1, e2);
}
