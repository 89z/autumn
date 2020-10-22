fn main() {
   // example 1
   let n1 = "10".parse().unwrap_or(0);
   // example 2
   let n2 = "AB".parse().unwrap_or(0);
   // print
   println!("{}", n1 == 10 && n2 == 0);
}
