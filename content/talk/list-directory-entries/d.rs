use std::fs;
fn main() {
   match fs::read_dir(".") {
      Err(e) => println!("{}", e),
      Ok(readdir) => for r_dirent in readdir {
         match r_dirent {
            Err(e) => println!("{}", e),
            // std::fs::DirEntry doesn't implement std::fmt::Display
            Ok(dirent) => println!("{}", dirent)
         }
      }
   }
}
