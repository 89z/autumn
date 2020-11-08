fn main() {
   // example 1
   let e: Result<u8, u8> = Ok(10);
   let n1 = e.unwrap_or(0);
   // example 2
   let e: Result<u8, u8> = Err(10);
   let n2 = e.unwrap_or(0);
   // print
   println!("{}", n1 == 10 && n2 == 0);
}
