fn main() {
   #[derive(Debug)]
   struct Pref {
      k: String,
      n: u8,
   }
   let o1 = Pref { k: "Sun".to_string(), n: 10 };
   println!("{:?}", o1);
}
