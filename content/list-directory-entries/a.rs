use std::fs;
fn main() {
   match fs::read_dir(".") {
      Err(e1) => println!("Err {:?}", e1),
      Ok(o1) => for o2 in o1 {
         println!("{:?}", o2);
      }
   }
}
