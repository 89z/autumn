use std::collections::HashMap;

fn main() -> Result<(), serde_json::Error> {
   let mut m = HashMap::new();
   m.insert("month", 12);
   m.insert("day", 31);
   let s = serde_json::to_string_pretty(&m)?;
   println!("{}", s);
   Ok(())
}
