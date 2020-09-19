fn main() {
   let s = "May";
   // example 1
   let s1 = &s[..2];
   // example 2
   let s2 = s.get(..2);
   // print
   println!("{}", s1 == "Ma" && s2 == Some("Ma"));
}
