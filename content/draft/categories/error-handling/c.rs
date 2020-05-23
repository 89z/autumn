use std::error::Error;
fn main() -> Result<(), Box<dyn Error>> {
   let s1 = "Sunday Monday";
   let mut a1 = s1.split_whitespace();
   // test 1
   let s2 = a1.nth(1).ok_or("None")?;
   println!("{}", s2);
   // test 2
   let s3 = a1.nth(2).ok_or("None")?;
   println!("{}", s3);
   // test 3
   println!("Sunday");
   Ok(())
}
