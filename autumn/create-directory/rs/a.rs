use std::fs;

fn main() {
   // example 1
   let e = fs::create_dir("March");
   println!("{:?}", e);

   // example 2
   let e = fs::create_dir_all("May/June");
   println!("{:?}", e);
}
