fn main() {
   let b = match 1 + 2 {
      5 => false,
      4 | 3 => true,
      n => n < 3
   };

   println!("{}", b);
}
