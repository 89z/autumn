fn main() {
   // example 1
   let n1 = Some(10).unwrap_or(0);
   // example 2
   let n2 = None.unwrap_or(0);
   // print
   println!("{}", n1 == 10 && n2 == 0);
}
