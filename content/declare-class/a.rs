fn main() {

   #[derive(Debug)]
   struct Date {
      year: u16,
      month: Option<u8>,
   }

   let o = Date {year: 2019, month: None};
   println!("{:?}", o);
}
