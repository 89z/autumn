fn main() {
   // example 1
   let s1: &str = "one\\two";
   // example 2
   let s2 = "one\\two";
   // example 3
   let s3 = r"one\two";
   // example 4
   let s4 = String::new();
   // example 5
   let s5 = String::from("one\\two");
   // print
   println!("{},{},{},{},{}", s1, s2, s3, s4, s5);
}
