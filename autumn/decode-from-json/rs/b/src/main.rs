use json;

fn main() -> Result<(), json::JsonError> {
   let s = r#"
{"month": 12, "day": 31}
"#;
   let m = json::parse(s)?;
   let n = &m["day"];
   println!("{}", n == 31);
   Ok(())
}
