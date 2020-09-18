fn main() {
   let a = vec!["May", "June"];
   // example A
   let sa = a.iter().fold(String::new(), |sc, sd| sc + sd);
   // example B
   let f = |sc: String, sd: &&str| sc + sd;
   let sb = a.iter().fold(String::new(), f);
   // print
   println!("{}", sa == "MayJune" && sb == "MayJune");
}
