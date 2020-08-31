use std::path::Path;
fn main() {
   let s1 = "index.md";
   let o1 = Path::new(s1);
   let s2 = o1.extension();
   println!("{:?}", s2);
}
