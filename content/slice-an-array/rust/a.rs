fn main() {
   let a = vec!["May", "June", "July"];
   // example 1
   let s1 = a[0];
   // example 2
   let s2 = a[a.len() - 1];
   // print
   println!("{}", s1 == "May" && s2 == "July");
}
