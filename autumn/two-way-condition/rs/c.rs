fn main() {
   // example 1
   let n = Some(10).unwrap_or(0);
   println!("{}", n == 10);
   // example 2
   let n = None.unwrap_or(0);
   println!("{}", n == 0);
}
