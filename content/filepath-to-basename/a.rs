use std::path::Path;
fn main() -> Result<(), String> {
   let o = Path::new("C:\\Windows\\write.exe");
   // example 1
   if let Some(s) = o.file_name() {
      println!("{:?}", s)
   } else {
      return Err("file_name".into())
   }
   // example 2
   if let Some(s) = o.file_stem() {
      println!("{:?}", s)
   } else {
      return Err("file_stem".into())
   }
   // return
   Ok(())
}
