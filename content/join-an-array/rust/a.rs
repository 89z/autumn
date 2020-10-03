fn main() {
   let a = vec!["May", "June"];
   // example 1
   let s1 = a.concat();
   // example 2
   let s2 = a.join(",");
   // print
   println!("{}", s1 == "MayJune" && s2 == "May,June");
}
