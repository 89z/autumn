fn main() {
   let a = String::from("March").into_bytes();
   let s = String::from_utf8_lossy(&a);
   println!("{}", s);
}
