fn main() {

   // example 1
   let r: Result<u8, u8> = Ok(10);
   let r = r.map_err(|e|
      e.to_string()
   );
   println!("{}", r == Ok(10));

   // example 2
   let r: Result<u8, u8> = Err(10);
   let r = r.map_err(|e|
      e.to_string()
   );
   println!("{}", r == Err(String::from("10")));

}
