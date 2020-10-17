fn main() {
   let s = "March";
   // example 1
   let s1 = &s[.. 2];
   // example 2
   let s2 = &s[2 ..];
   // example 3
   let s3 = &s[2 .. 3];
   // print
   println!("{}", s1 == "Ma" && s2 == "rch" && s3 == "r");
}
