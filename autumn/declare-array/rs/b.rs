fn main() {
   // example 1
   let a1: &[u8] = &[10, 11];
   // example 2
   let a2 = &[10, 11];
   // example 3
   let a3 = &[
      &[10, 11], &[12, 13]
   ];
   // print
   println!("{:?} {:?} {:?}", a1, a2, a3);
}
