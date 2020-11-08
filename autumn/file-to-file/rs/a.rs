use std::fs;

fn main() {
   let e = fs::copy("a.txt", "b.txt");
   println!("{:?}", e);
}
