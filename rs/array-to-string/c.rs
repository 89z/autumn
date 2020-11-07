use std::string;

fn main() -> Result<(), string::FromUtf8Error> {
   let a = String::from("March").into_bytes();
   let s = String::from_utf8(a)?;
   println!("{}", s);
   Ok(())
}
