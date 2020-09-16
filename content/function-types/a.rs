// example 1
fn f(n: u8, n2: u8) -> bool {
   return n > n2;
}
// example 2
fn f2(n: u8, n2: u8) -> bool {
   n > n2
}
// print
fn main() {
   let b = f(9, 8);
   let b2 = f2(9, 8);
   println!("{} {}", b, b2);
}
