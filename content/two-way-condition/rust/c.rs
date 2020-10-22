fn main() {
   // example 1
   let n1 = Some(10).unwrap_or_default();
   // example 2
   let n2: u8 = None.unwrap_or_default();
   // print
   println!("{}", n1 == 10 && n2 == 0);
}
