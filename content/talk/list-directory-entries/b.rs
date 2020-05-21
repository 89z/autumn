use std::ffi::OsStr;
fn main() {
   let s1 = OsStr::new("Sunday");
   // std::ffi::OsStr doesn't implement std::fmt::Display
   println!("{}", s1);
}
