fn main() {
   let n = 10;
   let s;
   if n > 0 {
      s = "+";
   } else if n < 0 {
      s = "-";
   } else {
      s = "zero";
   }
   println!("{}", s == "+");
}
