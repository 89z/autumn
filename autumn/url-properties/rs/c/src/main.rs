use {
   std::collections::HashMap,
   url::Url
};

fn query(u: Url) -> HashMap<String, String> {
   u.query_pairs().into_owned().collect()
}

fn main() -> Result<(), url::ParseError> {
   let u = Url::parse("http://docs.rs?west=left&east=right")?;
   let m = query(u);
   println!("{:?}", m);
   Ok(())
}
