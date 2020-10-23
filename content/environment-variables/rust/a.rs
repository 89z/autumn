use std::env;

fn main() {
   let r = env::var("USERPROFILE");
   let s = r.unwrap_or_default();
   println!("{}", s == r"C:\Users\Steven");
}
