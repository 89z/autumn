fn main() {
   let s = "May";
   let s1 = s.get(1 .. 2);
   println!("{}", s1 == Some("a"));
}
