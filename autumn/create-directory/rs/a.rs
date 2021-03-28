use std::fs;

fn main() {
   // example 1
   let e = fs::create_dir("north");
   println!("{:?}", e);

   // example 2
   let e = fs::create_dir_all("west/east");
   println!("{:?}", e);
}
