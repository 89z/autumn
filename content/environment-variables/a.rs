use std::env;

fn main() -> Result<(), env::VarError> {
   // example 1
   let s1 = env::var("USERPROFILE")?;
   // example 2
   let s2 = env!("USERPROFILE");
   // print
   println!("{}", s1 == r"C:\Users\Steven" && s2 == r"C:\Users\Steven");
   Ok(())
}
