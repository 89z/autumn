fn main() {
   let n = 10;
   let s = match n {
      12 => "all",
      11 | 10 => "some",
      _ => "none"
   };
   println!("{}", s == "some");
}
