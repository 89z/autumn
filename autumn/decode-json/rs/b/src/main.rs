use json;

fn main() -> Result<(), json::JsonError> {
   let s = r#"
{"U2": {"Boy": ["Twilight", "I Will Follow"]}}
"#;
   let m = json::parse(s)?;
   let s = &m["U2"]["Boy"][0];
   println!("{}", s == "Twilight");
   Ok(())
}
