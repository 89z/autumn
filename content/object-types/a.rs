fn main() {
   #[derive(Debug)]
   struct Pref {
      b: Option<bool>,
      k: Option<String>,
      n: Option<u8>
   }
   let o1 = Pref {
      b: Some(true),
      k: None,
      n: Some(10)
   };
   dbg!(o1);
}
