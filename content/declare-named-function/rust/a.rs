// example 1
fn f1(s: &str, c: &str) -> bool {
   return &s[.. 1] == c;
}
// example 2
fn f2(s: &str, c: &str) -> bool {
   &s[.. 1] == c
}
// print
fn main() {
   let b1 = f1("June", "J");
   let b2 = f2("June", "J");
   println!("{}", b1 && b2);
}
