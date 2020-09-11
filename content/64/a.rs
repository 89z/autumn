use std::path::Path;

fn main() {
   let o = Path::new("C:\\Windows\\write.exe");
   // example 1
   let o2 = o.file_name();
   // example 2
   let o3 = o.file_stem();
   // print
   println!("{:?} {:?}", o2, o3);
}
