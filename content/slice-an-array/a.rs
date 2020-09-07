fn main() {
   let a = vec![10, 11, 12];
   // example 1
   let n = a[0];
   // example 2
   let n2 = a[a.len() - 1];
   // print
   println!("{} {}", n == 10, n2 == 12);
}
