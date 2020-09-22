use std::env;

fn main() -> Result<(), env::VarError> {
   let s = env::var("USERPROFILE")?;
   println!("{}", s == r"C:\Users\Steven");
   Ok(())
}
