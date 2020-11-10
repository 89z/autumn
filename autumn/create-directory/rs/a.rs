use std::fs;

fn main() {
   // example 1
   let e1 = fs::create_dir("March");
   // example 2
   let e2 = fs::create_dir_all("May/June");
   // print
   println!("{:?} {:?}", e1, e2);
}
