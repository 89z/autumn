use std::path::PathBuf;

fn main() {
   let mut p = PathBuf::from("south");
   p.push("north");
   println!("{:?}", p);
}
