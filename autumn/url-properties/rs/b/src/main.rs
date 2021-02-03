use url::Url;

fn main() -> Result<(), url::ParseError> {
   let u = Url::parse("http://docs.rs/url")?;
   let s = u.host_str().unwrap_or_default();
   println!("{}", s);
   Ok(())
}
