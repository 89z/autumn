fn main() {
   let n = 10;
   let s = if n > 0 {
      "+"
   } else if n < 0 {
      "-"
   } else {
      "zero"
   };
   println!("{}", s == "+");
}
