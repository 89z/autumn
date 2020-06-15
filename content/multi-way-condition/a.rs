fn main() {
   let s = "Sun";
   let n = match s {
      "Fri" => 1,
      "Sat" | "Sun" => 2,
      _ => 0
   };
   println!("{}", n);
}
