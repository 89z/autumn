use std::fs;

fn main() {
   let e = fs::remove_file("a.txt");
   println!("{:?}", e);
}
