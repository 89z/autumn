use std::fs;
fn main() {
   match fs::read_dir(".") {
      Err(e) => println!("{}", e),
      Ok(readdir) => for r_dirent in readdir {
         match r_dirent {
            Err(e) => println!("{}", e),
            Ok(dirent) => {
               let name = dirent.file_name();
               // std::ffi::OsString` doesn't implement `std::fmt::Display
               println!("{}", name);
            }
         }
      }
   }
}
