use {
   std::collections::HashMap,
   url::Url
};

fn query(u: Url) -> HashMap<String, String> {
   u.query_pairs().into_owned().collect()
}

fn main() -> Result<(), url::ParseError> {
   let u = Url::parse("http://github.com?month=May&day=Friday")?;
   let q = query(u);
   println!("{:?}", q);
   Ok(())
}
