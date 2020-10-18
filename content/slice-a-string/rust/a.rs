fn main() {
   let s = "March";
   let s2 = s.chars().nth(2);
   println!("{}", s2 == Some('r'));
}
