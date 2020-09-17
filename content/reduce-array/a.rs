fn main() {
   let a = vec!["May", "June"];
   // example 1
   let s = a.iter().fold(String::new(), |sa, sc| sa + sc);
   // example 2
   let f = |sa: String, sc: &&str| sa + sc;
   let s2 = a.iter().fold(String::new(), f);
   // print
   println!("{}", s == "MayJune" && s2 == "MayJune");
}
