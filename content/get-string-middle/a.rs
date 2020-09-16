fn main() {
   let s = "June";
   // example 1
   let s2 = &s[1 .. 3];
   // example 2
   let s3 = s.get(1 .. 3);
   // print
   println!("{}", s2 == "un" && s3 == Some("un"));
}
