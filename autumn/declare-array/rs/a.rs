fn main() {
   // example 1
   let a: Vec<u8> = vec![];
   println!("{:?}", a);

   // example 2
   let a = vec![10, 11];
   println!("{:?}", a);

   // example 3
   let a = vec![
      vec![10, 11], vec![12, 13]
   ];
   println!("{:?}", a);
}
