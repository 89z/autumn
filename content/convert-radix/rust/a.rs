use std::char::from_digit;

fn encode(n: u32, r: u32) -> String {
   let c = from_digit(n % r, r).unwrap_or('!');
   return match n / r {
      0 => String::new(),
      n => encode(n, r)
   } + &String::from(c);
}

fn main() {
   let n = 1577858399;
   let s = encode(n, 36);
   println!("{}", s == "q3ezbz");
}
