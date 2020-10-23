use std::fs;

fn main() {
   let r = fs::remove_file("a.txt");
   println!("{:?}", r);
}
