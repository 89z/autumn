use tinyjson::{
   JsonParseError,
   JsonValue
};

fn main() -> Result<(), JsonParseError> {
   let s = r#"
{"month": 12, "day": 31}
"#;
   let m: JsonValue = s.parse()?;
   let n = &m["day"];
   println!("{}", n == &JsonValue::Number(31.0));
   Ok(())
}
