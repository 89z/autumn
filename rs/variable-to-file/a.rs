use std::fs;

fn main() {
   let e = fs::write("a.txt", "May\n");
   println!("{:?}", e);
}
