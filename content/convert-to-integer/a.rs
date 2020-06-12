fn main() {
   let s = "10";
   let n = s.parse::<u8>();
   println!("{}", n == Ok(10));
}
