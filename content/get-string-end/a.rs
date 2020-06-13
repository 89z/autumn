fn main() {
   let s1 = "Sunday";
   let s2 = &s1[1..];
   println!("{}", s2 == "unday");
}
