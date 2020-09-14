fn main() {
   let n = 9;
   // example 1
   let s = format!("{}", n);
   // example 2
   let s2 = n.to_string();
   // print
   println!("{} {}", s == "9", s2 == "9");
}
