use tinyjson::JsonValue;

fn main() -> Result<(), String> {
   let s = r#"
{"month": 12, "day": 31}
"#;
   let o: JsonValue = match s.parse() {
      Ok(v) => v,
      Err(v) => Err(format!("{}", v))?
   };
   let n: f64 = match o["day"].get() {
      Some(v) => *v,
      None => Err("day")?
   };
   println!("{}", n == 31.0);
   Ok(())
}
