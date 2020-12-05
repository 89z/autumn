use chrono::{
   TimeZone,
   Utc
};

fn main() {
   let old_o = Utc.ymd(2018, 12, 31);
   let new_o = Utc.ymd(2019, 12, 31);
   let dur_o = new_o.signed_duration_since(old_o);
   println!("{}", dur_o);
}
