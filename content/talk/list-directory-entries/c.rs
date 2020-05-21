use std::ffi::OsString;
fn main() {
   let s1 = OsString::from("Sunday");
   // std::ffi::OsString doesn't implement std::fmt::Display
   println!("{}", s1);
}
