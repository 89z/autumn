use std::collections::HashMap;

trait Hash {
   fn to_map(&self) -> HashMap<&str, u16>;
}

impl Hash for [(&str, u16)] {
   fn to_map(&self) -> HashMap<&str, u16> {
      self.iter().cloned().collect()
   }
}

fn main() {
   let m = [("month", 12), ("day", 31)].to_map();
   println!("{:?}", m)
}
