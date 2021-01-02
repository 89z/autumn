use chrono::{
   DateTime,
   offset
};

fn main() -> Result<(), chrono::ParseError> {
   let s = "2020-12-31T23:59:59Z";
   let o: DateTime<offset::Utc> = s.parse()?;
   println!("{}", o);
   Ok(())
}
