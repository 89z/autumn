use std::error::Error;
use std::num::ParseIntError;
use csv::{Reader, Writer};

fn main() -> Result<(), Box<dyn Error>> {
   let mut reader = Reader::from_path("data.csv")?;
   let mut writer = Writer::from_path("output.csv")?;
   let mut headers = reader.headers()?.clone();
   headers.push_field("SUM");
   writer.write_record(headers.iter())?;
   for row in reader.records() {
      let mut row = row?;
      let sum: Result<_, ParseIntError> = row.iter().try_fold(0, |accum, s| {
         Ok(accum + s.parse::<i64>()?)
      });
      row.push_field(&sum?.to_string());
      writer.write_record(row.iter())?;
   }
   writer.flush()?;
   Ok(())
}
