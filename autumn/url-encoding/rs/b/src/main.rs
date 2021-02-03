use url::Url;

fn main() -> Result<(), url::ParseError> {
   let s = "http://docs.rs";
   let u = Url::parse(s)?;
   println!("{}", u);
   Ok(())
}
