use std::fs;
fn main() {
   // example 1
   let e1 = fs::create_dir("Sunday");
   println!("{:?}", e1);
   // example 2
   let e2 = fs::create_dir_all("Sunday/Monday");
   println!("{:?}", e2);
}
