// example 1
fn f1(s1: &str) -> usize {
   return s1.len();
}
// example 2
fn f2(s1: &str) -> usize {
   s1.len()
}
// print
fn main() {
   let n1 = f1("Sunday");
   let n2 = f2("Sunday");
   dbg!(n1, n2);
}
