use chrono::{
   TimeZone,
   Utc
};

fn main() -> Result<(), chrono::ParseError> {
   let s = "2020-05-04 03:02:01";
   let o  = Utc.datetime_from_str(s, "%Y-%m-%d %H:%M:%S")?;
   println!("{:?}", o);
   Ok(())
}
