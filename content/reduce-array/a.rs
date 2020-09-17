fn main() {
   let a = vec!["May", "June"];
   let s = a.iter().fold(String::new(), |sa, sc| sa + sc);
   println!("{}", s == "MayJune");
}
