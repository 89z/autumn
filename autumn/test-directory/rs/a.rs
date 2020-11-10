use std::path::Path;

fn main() {
   let o = Path::new(r"C:\Users");
   let b = o.is_dir();
   println!("{}", b);
}
