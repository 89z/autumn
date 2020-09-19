fn main() {
   let a = vec!["May", "June"];
   // example 1
   let s1 = a.iter().fold(String::new(), |s, s1| s + s1);
   // example 2
   let f = |s: String, s2: &&str| s + s2;
   let s2 = a.iter().fold(String::new(), f);
   // print
   println!("{}", s1 == "MayJune" && s2 == "MayJune");
}
