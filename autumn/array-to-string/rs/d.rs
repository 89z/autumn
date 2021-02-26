fn main() {
   let y = b"March".to_vec();
   let s = String::from_utf8_lossy(&y);
   println!("{}", s);
}
