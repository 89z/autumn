fn main() {
   let a = vec!["May", "June"];
   // example 1
   let s = a.concat();
   // example 2
   let s2 = a.join(",");
   // print
   println!("{}", s == "MayJune" && s2 == "May,June");
}
