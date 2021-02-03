use url::Url;

fn main() -> Result<(), url::ParseError> {
   let u = Url::parse("http://docs.rs/url")?;
   let s = u.path();
   println!("{}", s);
   Ok(())
}
