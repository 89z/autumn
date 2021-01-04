use chrono::{
   TimeZone,
   Utc
};

fn main() {
   let t = Utc.ymd(2020, 5, 4);
   let u = Utc::today();
   let d = u.signed_duration_since(t);
   println!("{}", d);
}
