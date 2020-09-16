fn main() {
   let s = "📗";
   // example 1
   let n = s.len();
   // example 2
   let n2 = s.chars().count();
   // print
   println!("{}", n == 4 && n2 == 1);
}
