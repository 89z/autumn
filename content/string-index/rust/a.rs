fn main() {
   let s = "March";
   let n = s.chars().position(|c| c == 'a');
   println!("{}", n == Some(1));
}
