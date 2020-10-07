fn main() {
   let s = "May";
   let s1 = s.chars().nth(1);
   println!("{}", s1 == Some('a'));
}
