fn main() {

   // example 1
   let f = |d: u8, e: u8| -> u8 {
      return d + e
   };
   println!("{}", f(4, 5) == 9);

   // example 2
   let f = |d, e| -> i32 {
      return d + e
   };
   println!("{}", f(4, 5) == 9);

   // example 3
   let f = |d, e| d + e;
   println!("{}", f(4, 5) == 9);

}
