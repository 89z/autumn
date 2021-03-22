fn main() {
   let a = ["May", "June"];

   // example 1
   let s = a.iter().fold(String::new(), |s, t| s + t);
   println!("{}", s == "MayJune");

   // example 2
   let n = a.iter().fold(0, |n, s| n + s.len());
   println!("{}", n == 7);
}
