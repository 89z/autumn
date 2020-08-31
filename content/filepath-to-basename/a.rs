use std::error::Error;
use std::path::Path;
fn main() -> Result<(), Box<dyn Error>> {
   let s1 = "C:\\Windows\\write.exe";
   let o1 = Path::new(s1);
   let s2 = o1.file_stem().ok_or("file_stem")?;
   println!("{}", s2 == "write");
   Ok(())
}
