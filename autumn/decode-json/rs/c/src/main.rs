use serde_json::{
   Error,
   Value
};

fn main() -> Result<(), Error> {
   let s = r#"
{"U2": {"Boy": ["Twilight", "I Will Follow"]}}
"#;
   let m: Value = serde_json::from_str(s)?;
   let s = &m["U2"]["Boy"][0];
   println!("{}", s == "Twilight");
   Ok(())
}
