fn main() {
   // example 1
   let n = Some(10).unwrap_or_default();
   println!("{}", n == 10);
   // example 2
   let n: u8 = None.unwrap_or_default();
   println!("{}", n == 0);
}
