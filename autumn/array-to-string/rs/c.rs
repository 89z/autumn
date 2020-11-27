use std::str;

fn main() -> Result<(), str::Utf8Error> {
   let y = b"March";
   let s = str::from_utf8(y)?;
   println!("{}", s);
   Ok(())
}
