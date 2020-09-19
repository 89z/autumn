use std::path::Path;

fn main() {
   let o = Path::new("C:\\Windows\\notepad.exe");
   // example 1
   let o1 = o.file_name();
   // example 2
   let o2 = o.file_stem();
   // print
   println!("{:?} {:?}", o1, o2);
}
