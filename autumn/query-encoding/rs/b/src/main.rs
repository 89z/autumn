use std::collections::HashMap;

fn decode(s: &str) -> HashMap<String, String> {
   let a = s.as_bytes();
   form_urlencoded::parse(a).into_owned().collect()
}

fn main() {
   let s = "month=May&day=Friday";
   let m = decode(s);
   println!("{:?}", m);
}
