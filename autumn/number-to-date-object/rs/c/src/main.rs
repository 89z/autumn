use chrono::{
   TimeZone,
   Utc
};

fn main() {
   // example 1
   let d = TimeZone::ymd(&Utc, 2020, 12, 31);
   println!("{}", d);
   // example 2
   let d = Utc.ymd(2020, 12, 31);
   println!("{}", d);
}
