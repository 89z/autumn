use std::path::Path;

fn main() {
   let p = Path::new("south").join("north");
   println!("{:?}", p);
}
