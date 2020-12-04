use chrono::prelude::*;

fn main() {
   let z  = Utc.datetime_from_str("2014-11-28 12:00:09", "%Y-%m-%d %H:%M:%S");
   println!("{:?}", z);
}
