fn main() {
   let s = "Sunday";
   // example 1
   let s1 = &s[1..3];
   // example 2
   let s2 = s.get(1..3);
   // print
   println!("{} {}", s1 == "un", s2 == Some("un"));
}
