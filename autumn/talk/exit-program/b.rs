use std::char;

fn main() -> Result<(), u32> {
   let n = 10;
   let c = char::from_digit(n, 16).ok_or(n)?;
   println!("{}", c);
   Ok(())
}
