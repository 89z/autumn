fn main() {

   // example 1
   let a = vec!["May", "June"];
   let s = a.join(",");

   // example 2
   let a = vec![10, 11];
   let mut s2 = String::new();
   for n in a.iter() {
      if s2 != "" {
         s2 += ",";
      }
      s2 += &n.to_string();
   }

   // print
   println!("{} {}", s == "May,June", s2 == "10,11");
}
