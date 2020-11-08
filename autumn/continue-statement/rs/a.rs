fn main() {
   for n in 0 .. 10 {
      if n % 3 == 0 {
         continue;
      }
      println!("{}", n);
   }
}
