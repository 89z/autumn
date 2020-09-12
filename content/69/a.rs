use std::path::Path;

fn main() {
   let b = Path::new("index.md").is_file();
   println!("{}", b);
}
