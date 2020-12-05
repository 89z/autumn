use chrono::prelude::*;
use chrono::Duration;

fn main() {
   assert_eq!(Utc.ymd(1970, 1, 1).and_hms(0, 0, 0) + Duration::seconds(1_000_000_000), Utc.ymd(2001, 9, 9).and_hms(1, 46, 40));
   assert_eq!(Utc.ymd(1970, 1, 1).and_hms(0, 0, 0) - Duration::seconds(1_000_000_000), Utc.ymd(1938, 4, 24).and_hms(22, 13, 20));
}
