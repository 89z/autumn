fn main() {
   let mut s = String::from("East");
   s.insert_str(0, "West");
   println!("{}", s);
}
