// example 1
fn f(s: &str) -> usize {
   return s.len();
}
// example 2
fn f2(s: &str) -> usize {
   s.len()
}
// print
fn main() {
   let n = f("May");
   let n2 = f2("May");
   println!("{} {}", n == 3, n2 == 3);
}
