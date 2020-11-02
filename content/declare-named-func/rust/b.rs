trait Comp {
   fn greater(self, u8) -> bool;
}

impl Comp for u8 {
   fn greater(self, n: u8) -> bool {
      self > n
   }
}

fn main() {
   let n = 10;
   let b = n.greater(9);
   println!("{}", b);
}
