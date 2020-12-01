use tinyjson::JsonValue;

fn main() -> Result<(), String> {
   let s = r#"
{"U2": {"Boy": ["Twilight", "I Will Follow"]}}
"#;
   let m: JsonValue = match s.parse() {
      Ok(v) => v,
      Err(v) => Err(format!("{}", v))?
   };
   let s = match &m["U2"]["Boy"][0] {
      JsonValue::String(v) => v,
      v => Err(format!("{:?}", v))?
   };
   println!("{}", s == "Twilight");
   Ok(())
}
