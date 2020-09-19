use std::path::Path;

fn main() {
   let o = Path::new("C:\\Windows\\notepad.exe");
   let s = o.parent();
   println!("{:?}", s);
}
