use std::fs;

fn main() {
   let e = fs::write("a.txt", "March\n");
   println!("{:?}", e);
}
