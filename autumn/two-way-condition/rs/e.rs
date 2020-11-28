fn main() {
   // example 1
   let r1: Result<u8, u8> = Ok(10);
   let r1 = r1.map_err(|e|
      e.to_string()
   );
   // example 2
   let r2: Result<u8, u8> = Err(10);
   let r2 = r2.map_err(|e|
      e.to_string()
   );
   // print
   println!("{}", r1 == Ok(10) && r2 == Err(String::from("10")));
}
