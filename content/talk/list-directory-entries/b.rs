use std::fs;
fn main() {
   match fs::read_dir(".") {
      Err(e) => println!("{}", e),
      Ok(readdir) => for r_dirent in readdir {
         match r_dirent {
            Err(e) => println!("{}", e),
            Ok(dirent) => {
               match dirent.path().to_str() {
                  None => println!("to_str"),
                  Some(s1) => println!("{}", s1)
               }
            }
         }
      }
   }
}
