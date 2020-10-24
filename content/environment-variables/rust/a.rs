use std::env;

fn main() {
   let e = env::var("USERPROFILE");
   let s = e.unwrap_or_default();
   println!("{}", s == r"C:\Users\Steven");
}
