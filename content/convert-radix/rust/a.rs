use std::char;

fn encode(n: u32, r: u32) -> String {
   if n == 0 {
      return String::new();
   }
   let c = char::from_digit(n % r, r).unwrap_or_default();
   encode(n / r, r) + &String::from(c)
}

fn main() {
   let n = 1577858399;
   let s = encode(n, 36);
   println!("{}", s == "q3ezbz");
}
