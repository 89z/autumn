fn main() {
   let s = "📗";
   // example 1
   let n1 = s.len();
   // example 2
   let n2 = s.chars().count();
   // print
   println!("{}", n1 == 4 && n2 == 1);
}
