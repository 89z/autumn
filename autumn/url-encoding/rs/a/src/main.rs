use url::Url;

fn main() -> Result<(), url::ParseError> {
   let u = Url::parse("http://docs.rs")?;
   let s = u.into_string();
   println!("{}", s);
   Ok(())
}
