use data_encoding::{BASE64, DecodeError};

fn main() -> Result<(), DecodeError> {
   let s = b"CgsMDQ==";
   let t = BASE64.decode(s)?;
   println!("{:?}", t);
   Ok(())
}
