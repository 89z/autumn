use std::path::Path;

fn main() {
   let p = Path::new(r"C:\Windows\notepad.exe");
   let s = p.file_stem().unwrap_or_default();
   println!("{}", s == "notepad");
}
