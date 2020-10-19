fn main() {
   let a = ["May", "June"];
   // example 1
   let s = a.iter().fold(String::new(), |s, s1| s + s1);
   // example 2
   let n = a.iter().fold(0, |n, s| n + s.len());
   // print
   println!("{}", s == "MayJune" && n == 7);
}
