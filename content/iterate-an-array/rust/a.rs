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
   // example 3
   for (n, s) in a.iter().enumerate() {
      println!("{} {}", n, s);
   }
}
