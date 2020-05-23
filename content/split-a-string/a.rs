fn main() {
   let s1 = "Sunday Monday";
   let mut o1 = s1.split_whitespace();
   let o2 = o1.nth(1);
   println!("{:?}", o2);
}
