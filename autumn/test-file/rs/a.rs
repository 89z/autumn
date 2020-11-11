use std::path::Path;

fn main() {
   let o = Path::new("a.rs");
   let b = o.is_file();
   println!("{}", b);
}
