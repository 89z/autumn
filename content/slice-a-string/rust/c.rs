fn main() {
   let s = "June";
   // example 1
   let s1 = &s[2 .. 3];
   // example 2
   let s2 = &s[2 ..= 2];
   // print
   println!("{}", s1 == "n" && s2 == "n");
}
