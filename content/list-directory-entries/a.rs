use std::fs;
fn main() {
   if let Ok(mut o1) = fs::read_dir(".") {
      while let Some(o2) = o1.next() {
         if let Ok(o3) = o2 {
            println!("{:?}", o3);
         }
      }
   }
}
