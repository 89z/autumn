fn main() {
   let s = "March";
   // example 1
   let s1 = s.get(.. 2);
   // example 2
   let s2 = s.get(2 ..);
   // example 3
   let s3 = s.get(2 .. 3);
   // print
   println!("{}", s1 == Some("Ma") && s2 == Some("rch") && s3 == Some("r"));
}
