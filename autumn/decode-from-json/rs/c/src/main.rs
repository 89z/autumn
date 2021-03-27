use serde_json::{
   Error,
   Value
};

fn main() -> Result<(), Error> {
   let s = r#"
{"month": 12, "day": 31}
"#;
   let m: Value = serde_json::from_str(s)?;
   let n = &m["day"];
   println!("{}", n == 31);
   Ok(())
}
