fn main() {
   let s = "March";
   let s1 = s.chars().nth(2);
   println!("{}", s1 == Some('r'));
}
