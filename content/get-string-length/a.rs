fn main() {
   let s1 = "♠";
   // example 1
   let n1 = s1.len();
   // example 2
   let n2 = s1.chars().count();
   // print
   dbg!(n1 == 3, n2 == 1);
}
