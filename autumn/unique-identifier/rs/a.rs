use std::char::from_digit;

fn encode(mut n: u32, r: u32) -> Option<String> {
   let mut s = String::new();
   loop {
      if let Some(c) = from_digit(n % r, r) {
         s.insert(0, c)
      } else {
         return None
      }
      n /= r;
      if n == 0 {
         break
      }
   }
   Some(s)
}

fn main() {
   let s = encode(u32::MAX, 36).unwrap_or_default();
   println!("{}", s == "1z141z3");
}
