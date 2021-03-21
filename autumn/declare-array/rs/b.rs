fn main() {
   // example 1
   let a: &[u8] = &[10, 11];
   println!("{:?}", a);
   // example 2
   let a = &[10, 11];
   println!("{:?}", a);
   // example 3
   let a = &[
      &[10, 11], &[12, 13]
   ];
   println!("{:?}", a);
}
