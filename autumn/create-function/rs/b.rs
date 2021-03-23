trait Oper {
   fn add(self, u8) -> u8;
}

impl Oper for u8 {
   fn add(self, n: u8) -> u8 {
      self + n
   }
}

fn main() {
   let n = 7;
   let n = n.add(1).add(1);
   println!("{}", n == 9);
}
