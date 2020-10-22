use std::fs;

fn main() {
   let r = fs::write("a.txt", "May\n");
   if let Err(e) = r {
      println!("{}", e);
   }
}
