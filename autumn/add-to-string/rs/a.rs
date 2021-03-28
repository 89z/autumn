fn main() {
   let mut s = String::from("West");
   let t = String::from("East");
   s = s + &t;
   println!("{}", s);
}
