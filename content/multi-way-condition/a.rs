fn main() {
   let n = 1;
   let s = match n {
      3 => "all",
      2 | 1 => "some",
      _ => "none"
   };
   println!("{}", s == "some");
}
