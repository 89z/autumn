fn main() {
   let a = [0, 10, 20, 30, 40];

   // example 1
   let n = a[2];
   println!("{}", n);

   // example 2
   let b = &a[2 ..];
   println!("{:?}", b);

   // example 3
   let b = &a[2 .. 4];
   println!("{:?}", b);

   // example 4
   let b = &a[2 ..= 3];
   println!("{:?}", b);

}
