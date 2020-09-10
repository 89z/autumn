fn main() {

   #[derive(Debug)]
   struct Date {
      year: u16,
      month: Option<u8>,
      day: Option<u8>
   }

   let o = Date {year: 2020, month: None, day: None};
   println!("{:?}", o);
}
