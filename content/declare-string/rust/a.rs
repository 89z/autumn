fn main() {
   // example 1
   let s1 = String::from("one\\two");
   // example 2
   let s2 = String::from(r"one\two");
   // print
   println!("{} {}", s1, s2);
}
