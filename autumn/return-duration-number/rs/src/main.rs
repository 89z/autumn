use chrono::{
   TimeZone,
   Utc
};

fn main() {
   let old = Utc.ymd(2019, 12, 31);
   let new = Utc.ymd(2020, 12, 31);
   let dur = new.signed_duration_since(old);
   println!("{}", dur);
}
