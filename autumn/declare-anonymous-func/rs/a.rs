fn main() {

   // example 1
   let f = |d: u8, e: u8| -> u8 {
      return d + e
   };

   // example 2
   let g = |d, e| -> i32 {
      return d + e
   };

   // example 3
   let h = |d, e| d + e;

   // print
   println!("{}", f(4, 5) == 9);
   println!("{}", g(4, 5) == 9);
   println!("{}", h(4, 5) == 9);
}
