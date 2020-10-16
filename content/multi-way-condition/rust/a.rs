fn main() {
   let s = "minute";

   let n = match s {
      "hour" => 23,
      "minute" | "second" => 59,
      _ => -1
   };

   println!("{}", n == 59);
}
