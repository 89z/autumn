use std::collections::HashMap;

fn decode(s: &str) -> HashMap<String, String> {
   let b = s.as_bytes();
   form_urlencoded::parse(b).into_owned().collect()
}

fn main() {
   let s = "west=left&east=right";
   let m = decode(s);
   println!("{:?}", m);
}
