fn main() {
   let mut s = String::from("April");
   s.insert_str(0, "March");
   println!("{}", s);
}
