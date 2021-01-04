use chrono::{
   DateTime,
   offset
};

fn main() -> Result<(), chrono::ParseError> {
   let s = "2020-05-04T03:02:01Z";
   let o: DateTime<offset::Utc> = s.parse()?;
   println!("{}", o);
   Ok(())
}
