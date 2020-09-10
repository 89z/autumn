fn main() {
   let a = vec!["May", "June", "July"];
   // example 1
   let s = a[0];
   // example 2
   let s2 = a[a.len() - 1];
   // print
   println!("{} {}", s == "May", s2 == "July");
}
