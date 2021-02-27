fn main() {
   let a = b"March".to_vec();
   let s = String::from_utf8_lossy(&a);
   println!("{}", s);
}
