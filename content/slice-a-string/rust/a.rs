fn main() {
   let s = "June";
   let s2 = s.chars().nth(2);
   println!("{}", s2 == Some('n'));
}
