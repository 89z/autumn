fn main() {
   let mut s = String::from("March");
   let t = String::from("April");
   s += &t;
   println!("{}", s);
}
