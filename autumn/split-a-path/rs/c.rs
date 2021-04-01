use std::path::Path;

fn main() {
   let p = Path::new(r"C:\Windows\notepad.exe");
   let q = match p.parent() {
      Some(v) => v,
      None => Path::new(r"C:\")
   };
   println!("{:?}", q);
}
