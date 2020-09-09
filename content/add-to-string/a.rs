fn main() {
   // example 1
   let mut s = "May".to_string();
   s += "June";
   // example 2
   let mut s2 = "May".to_owned();
   s2 += "June";
   // print
   println!("{} {}", s == "MayJune", s2 == "MayJune");
}
