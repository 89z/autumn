use std::error::Error;

fn main() -> Result<(), Box<dyn Error>> {

   // value left
   let s = "100";
   let n: u32 = s.parse()?;
   println!("{}", n == 100);

   // value right
   let n: u32 = str::parse("100")?;
   println!("{}", n == 100);

   // type right
   let n = str::parse::<u32>("100")?;
   println!("{}", n == 100);

   // float
   let n: f32 = str::parse("99.9")?;
   println!("{}", n == 99.9);

   Ok(())
}
