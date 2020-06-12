fn main() {
   let s1 = "Sunday";
   // example 1
   let s2 = &s1[1..3];
   // example 2
   let s3 = s1.get(1..3);
   // print
   println!("{} {}", s2 == "un", s3 == Some("un"));
}
