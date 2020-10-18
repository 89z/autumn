fn main() {
   let s = "March";
   // example 1
   let s1 = s.get(2 .. 3);
   // example 2
   let s2 = s.get(2 ..= 2);
   // print
   println!("{}", s1 == Some("r") && s2 == Some("r"));
}
