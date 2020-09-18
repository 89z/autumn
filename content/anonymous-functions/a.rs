fn main() {
   // example 1
   let f1 = |n: u8, n2: u8| -> bool {
      return n > n2;
   };
   // example 2
   let f2 = |n, n2| {
      return n > n2;
   };
   // example 3
   let f3 = |n, n2| n > n2;
   // print
   let b1 = f1(9, 8);
   let b2 = f2(9, 8);
   let b3 = f3(9, 8);
   println!("{}", b1 && b2 && b3);
}
