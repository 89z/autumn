fn main() {
   // example 1
   let s1 = "May";
   // example 2
   let s2: &str = "May";
   // example 3
   let s3 = String::from("May");
   // example 4
   let s4: String = "May".into();
   // print
   println!("{}{}{}{}", s1, s2, s3, s4);
}
