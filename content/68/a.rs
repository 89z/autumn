use std::path::Path;
fn main() {
   let o = Path::new("C:\\Windows\\write.exe");
   let s = o.parent();
   println!("{:?}", s);
}
