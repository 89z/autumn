fn main() {
   let (f, i) = (7.5, 7);

   // example 1
   let n = f / 2.0;
   println!("{}", n == 3.75);

   // example 2
   let n = f as u8 / 2;
   println!("{}", n == 3);

   // example 3
   let n = i / 2;
   println!("{}", n == 3);

   // example 4
   let n = i as f32 / 2.0;
   println!("{}", n == 3.5);

}
