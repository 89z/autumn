use chrono::{
   DateTime,
   offset
};

fn main() -> Result<(), chrono::ParseError> {
   let s = "2020-12-31T01:02:31Z";
   let d: DateTime<offset::Utc> = s.parse()?;
   println!("{}", d);
   Ok(())
}
