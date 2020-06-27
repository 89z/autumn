fn main() {
   let n = 10;
   let s: &str;
   if n > 12 {
      s = "Tue";
   } else if n > 11 {
      s = "Mon";
   } else {
      s = "Sun";
   }
   println!("{}", s);
}
