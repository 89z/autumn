use curl::easy::Easy;
use std::io::{stdout, Write};
fn main() {
   let s1 = "http://speedtest.lax.hivelocity.net";
   let mut easy = Easy::new();
   easy.url(s1).unwrap();
   easy.write_function(|data| {
      stdout().write_all(data).unwrap();
      Ok(data.len())
   }).unwrap();
   easy.perform().unwrap();
}
