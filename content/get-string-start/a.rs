fn main() {
   let s1 = "Sunday";
   // example 1
   let s2 = &s1[..2];
   // example 2
   let s3 = s1.get(..2);
   // print
   println!("{} {}", s2 == "Su", s3 == Some("Su"));
}
