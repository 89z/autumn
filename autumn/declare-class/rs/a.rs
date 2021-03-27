struct Date { month: u8, day: u8 }

impl Date {
   fn new() -> Self {
      Self { month: 12, day: 31 }
   }
   fn year(&mut self) {
      self.month = 1;
      self.day = 1;
   }
}

fn main () {
   let mut d = Date::new();
   d.year();
   println!("{}", d.day == 1);
}
