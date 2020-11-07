struct Time {
   hours: u16
}

impl Time {
   fn duration(self, minutes: u16) -> u16 {
      self.hours * 60 + minutes
   }
}

fn main () {
   let o = Time{hours: 23};
   let n = o.duration(59);
   println!("{}", n == 1439);
}
