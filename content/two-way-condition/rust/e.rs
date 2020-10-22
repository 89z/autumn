fn main() {
   // example 1
   let n1: u8 = "10".parse().unwrap_or_default();
   // example 2
   let n2: u8 = "AB".parse().unwrap_or_default();
   // print
   println!("{}", n1 == 10 && n2 == 0);
}
