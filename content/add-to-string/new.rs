fn main() {
   let mut s = String::new();
   // example 1
   s.push_str("May");
   // example 2
   s += "June";
   // print
   println!("{}", s == "MayJune");
}
