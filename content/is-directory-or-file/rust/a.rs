use std::path::Path;

fn main() {
   let o = Path::new("index.md");
   let b = o.is_file();
   println!("{}", b);
}
