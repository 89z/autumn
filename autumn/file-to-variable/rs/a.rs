use std::fs;

fn main() -> Result<(), String> {
   let s = "a.rs";
   match fs::read_to_string(s) {
      Ok(v) => print!("{}", v),
      Err(v) => Err(format!("{} {}", s, v))?
   }
   Ok(())
}
