use std::io;
fn main() -> io::Result<()> {
   let mut s = String::new();
   io::stdin().read_line(&mut s)?;
   println!("{}", s);
   Ok(())
}
