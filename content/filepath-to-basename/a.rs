use std::error::Error;
use std::path::Path;
fn main() -> Result<(), Box<dyn Error>> {
   let s = "C:\\Windows\\write.exe";
   // example 1
   let s1 = Path::new(s).file_name().ok_or("file_name")?;
   // example 2
   let s2 = Path::new(s).file_stem().ok_or("file_stem")?;
   // print
   println!("{} {}", s1 == "write.exe", s2 == "write");
   Ok(())
}
