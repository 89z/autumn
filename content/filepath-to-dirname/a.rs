use std::path::Path;

fn main() {
   let o = Path::new(r"C:\Windows\notepad.exe");
   let s = o.parent();
   println!("{:?}", s);
}
