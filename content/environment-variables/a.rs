use std::env;

fn main() -> Result<(), env::VarError> {
   let s = env::var("BROWSER")?;
   println!("{}", s);
   Ok(())
}
