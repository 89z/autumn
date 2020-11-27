use std::string;

fn main() -> Result<(), string::FromUtf8Error> {
   let y = b"March".to_vec();
   let s = String::from_utf8(y)?;
   println!("{}", s);
   Ok(())
}
