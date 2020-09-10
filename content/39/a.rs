fn main() {
   let s = "May";
   // example 1
   let s2 = &s[..2];
   // example 2
   let s3 = s.get(..2);
   // print
   println!("{} {}", s2 == "Ma", s3 == Some("Ma"));
}
