fn main() {
   let a = ["M", "a", "r", "c", "h"];

   // example 1
   let o = a.get(2);
   println!("{:?}", o);

   // example 2
   let o = a.get(2 .. 4);
   println!("{:?}", o);

   // example 3
   let o = a.get(2 ..= 3);
   println!("{:?}", o);

   // example 4
   let o = a.get(2 ..);
   println!("{:?}", o);
}
