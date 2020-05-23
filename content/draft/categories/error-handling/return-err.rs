use std::collections::HashMap;
fn main() -> Result<(), ()> {
   let mut m1 = HashMap::new();
   m1.insert("Sunday", 10);
   // test 1
   let n1 = match m1.get("Sunday") {
      None => return Err(()),
      Some(n) => n
   };
   println!("{}", n1);
   // test 2
   let n2 = match m1.get("Monday") {
      None => return Err(()),
      Some(n) => n
   };
   println!("{}", n2);
   // test 3
   println!("Ok");
   Ok(())
}
