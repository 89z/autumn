fn main() {
   let mut s = String::from("May");
   // example 1
   s.push_str("June");
   // example 2
   s += "July";
   // print
   println!("{}", s == "MayJuneJuly");
}
