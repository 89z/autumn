fn main() {
   let mut s = String::from("March");
   let t = String::from("April");
   s = s + &t;
   println!("{}", s);
}
