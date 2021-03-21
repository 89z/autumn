fn main() {
   // example 1
   let e: Result<u8, u8> = Some(10).ok_or(1);
   println!("{}", e == Ok(10));
   // example 2
   let e: Result<u8, u8> = None.ok_or(1);
   println!("{}", e == Err(1));
}
