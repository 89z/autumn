fn main() {
   // example 1
   let e: Result<u8, u8> = Ok(10);
   let n = e.unwrap_or_default();
   println!("{}", n == 10);

   // example 2
   let e: Result<u8, u8> = Err(10);
   let n = e.unwrap_or_default();
   println!("{}", n == 0);
}
