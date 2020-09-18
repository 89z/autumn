fn main() {
   let a = vec!["May", "June", "July"];
   // example A
   let sa = a[0];
   // example B
   let sb = a[a.len() - 1];
   // print
   println!("{}", sa == "May" && sb == "July");
}
