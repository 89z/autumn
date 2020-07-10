use std::path::Path;
fn main() {
   let b = Path::new("index.md").exists();
   println!("{}", b);
}
