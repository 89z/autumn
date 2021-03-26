use std::path::Path;

fn main() {
   let p = Path::new("a.rs");
   let b = p.is_file();
   println!("{}", b);
}
