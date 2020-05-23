use std::collections::HashMap;
fn main() {
   let mut m1 = HashMap::new();
   m1.insert("Sunday", 10);
   let n1 = match m1.get("Sunday") {
      None => {
         println!("None");
         return;
      },
      Some(n) => n
   };
   println!("{}", n1);
}
