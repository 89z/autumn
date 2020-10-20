fn main() {
   // example 1
   let mut s1 = String::new();
   s1 += "May";
   // example 2
   let mut s2 = String::from("May");
   s2 += "June";
   // print
   println!("{}", s1 == "May" && s2 == "MayJune");
}
