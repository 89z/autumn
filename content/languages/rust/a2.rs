fn main() {
   #[derive(Debug)]
   struct Pref {
      k: String,
      n: u8,
      b: bool
   }
   // error[E0063]: missing field `b` in initializer of `main::Pref`
   let o1 = Pref { k: "Sun".to_string(), n: 10 };
   println!("{:?}", o1);
}
