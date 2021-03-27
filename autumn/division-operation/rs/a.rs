fn main() {

   // natural int
   let n = 7 / 2;
   println!("{}", n == 3);

   // natural float
   let n = 7.0 / 2.0;
   println!("{}", n == 3.5);

   // force int
   let n = 7.0 as u32 / 2;
   println!("{}", n == 3);

   // force float
   let n = 7 as f32 / 2.0;
   println!("{}", n == 3.5);

}
