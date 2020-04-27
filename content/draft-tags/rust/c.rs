fn main() {
   #[derive(Debug, Default)]
   struct Pref {
      k: String,
      n: u8,
      b: bool
   }
   let o1 = Pref { k: "Sun".to_string(), n: 10, ..Default::default() };
   println!("{:?}", o1);
}
