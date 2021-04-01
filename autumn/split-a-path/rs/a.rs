use std::path::Path;

fn main() {
   let p = Path::new("a.rs");
   let s = p.extension().unwrap_or_default();
   println!("{}", s == "rs");
}
