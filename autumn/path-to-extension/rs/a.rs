use std::path::Path;

fn main() {
   let p = Path::new("a.rs");
   let o = p.extension();
   println!("{:?}", o);
}
