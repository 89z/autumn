fn main() {
   let a = ["M", "a", "r", "c", "h"];

   // example 1
   let s = a[2];
   println!("{}", s);

   // example 2
   let b = &a[2 .. 4];
   println!("{:?}", b);

   // example 3
   let b = &a[2 ..= 3];
   println!("{:?}", b);

   // example 4
   let b = &a[2 ..];
   println!("{:?}", b);
}
