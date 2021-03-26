use std::path::Path;

fn main() {
   let p = Path::new(r"C:\Users");
   let b = p.is_dir();
   println!("{}", b);
}
