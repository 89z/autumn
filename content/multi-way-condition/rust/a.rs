fn main() {
   let s = match 46 / 10 {
      7 => format!("seven"),
      6 | 5 => format!("six or five"),
      n => format!("{}", n)
   };

   println!("{}", s);
}
