use form_urlencoded::Serializer;

fn encode(m: &[(&str, &str)]) -> String {
   let s = String::new();
   Serializer::new(s).extend_pairs(m).finish()
}

fn main() {
   let m = &[("west", "left"), ("east", "right")];
   let s = encode(m);
   println!("{}", s == "west=left&east=right");
}
