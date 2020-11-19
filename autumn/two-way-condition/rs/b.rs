fn main() {
   // example 1
   let e1: Result<u8, u8> = Some(10).ok_or(1);
   // example 2
   let e2: Result<u8, u8> = None.ok_or(1);
   // print
   println!("{}", e1 == Ok(10) && e2 == Err(1));
}
