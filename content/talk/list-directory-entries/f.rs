use std::fs;
fn main() {
   match fs::read_dir(".") {
      Err(e) => println!("{}", e),
      Ok(readdir) => for r_dirent in readdir {
         match r_dirent {
            Err(e) => println!("{}", e),
            Ok(dirent) => {
               let pathbuf = dirent.path();
               // std::path::PathBuf doesn't implement std::fmt::Display
               println!("{}", pathbuf);
            }
         }
      }
   }
}
