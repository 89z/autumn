fn main() {
   let s1 = "Sunday Monday";
   let mut a1 = s1.split_whitespace();
   let s2 = a1.nth(1);
   dbg!(s2);
}
