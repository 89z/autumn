fn main() {
   let a = [b'w', b'e', b's', b't'];
   let s = String::from_utf8_lossy(&a);
   println!("{}", s);
}
