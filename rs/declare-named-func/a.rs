// example 1
fn f1(n: u8, n1: u8) -> bool {
   return n > n1;
}
// example 2
fn f2(n: u8, n1: u8) -> bool {
   n > n1
}
// print
fn main() {
   let b1 = f1(9, 8);
   let b2 = f2(9, 8);
   println!("{}", b1 && b2);
}
