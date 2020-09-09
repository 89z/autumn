fn main() {
   let s = "May June";
   let a: Vec<&str> = s.split_whitespace().collect();
   println!("{:?}", a);
}
