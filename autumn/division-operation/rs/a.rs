fn main() {
   let (f, i) = (7.5, 7);
   // example 1
   let n1 = f / 2f32;
   // example 2
   let n2 = f as u8 / 2;
   // example 3
   let n3 = i / 2;
   // example 4
   let n4 = i as f32 / 2f32;
   // print
   println!("{}", n1 == 3.75 && n2 == 3 && n3 == 3 && n4 == 3.5);
}
