fn main() {
   // example 1
   let mut s1 = "Sunday".to_string();
   s1 += "Monday";
   // example 2
   let mut s2 = "Sunday".to_owned();
   s2 += "Monday";
   // print
   dbg!(s1, s2);
}
