fn main() {
   let a1 = ["Sun", "Mon"];
   // example 1
   for s1 in &a1 {
      dbg!(s1);
   }
   // example 2
   for s1 in a1.iter() {
      dbg!(s1);
   }
}
