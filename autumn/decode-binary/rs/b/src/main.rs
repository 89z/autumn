use data_encoding::{DecodeError, HEXUPPER};

fn main() -> Result<(), DecodeError> {
   let s = b"0A0B0C0D";
   let t = HEXUPPER.decode(s)?;
   println!("{:?}", t);
   Ok(())
}
