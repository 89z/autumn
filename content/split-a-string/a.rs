fn main() {
   let s1 = "Sunday Monday";
   let a1: Vec<&str> = s1.split_whitespace().collect();
   println!("{:?}", a1);
}
