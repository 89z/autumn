fn main() {
   let s = "Sunday Monday";
   let a: Vec<&str> = s.split_whitespace().collect();
   println!("{:?}", a);
}
