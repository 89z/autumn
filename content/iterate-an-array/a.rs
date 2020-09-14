fn main() {
   let a = ["May", "June"];
   // example 1
   for s in &a {
      println!("{}", s);
   }
   // example 2
   for s in a.iter() {
      println!("{}", s);
   }
}
