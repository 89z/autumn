use chrono::{
   TimeZone,
   Utc
};

fn main() -> Result<(), chrono::ParseError> {
   let s = "2020-12-31 01:02:31";
   let t  = Utc.datetime_from_str(s, "%Y-%m-%d %H:%M:%S")?;
   println!("{:?}", t);
   Ok(())
}
