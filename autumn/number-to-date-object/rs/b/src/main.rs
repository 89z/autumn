use chrono::{
   TimeZone,
   Utc
};

fn main() {
   let o = Utc.ymd(2020, 5, 4);
   println!("{}", o);
}
