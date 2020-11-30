use std::fs;

fn main() -> Result<(), String> {
   let o = fs::metadata("a.rs").map_err(|e| format!("{}", e))?;
   println!("{}", o.len());
   Ok(())
}
