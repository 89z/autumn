use std::fs;

fn main() {
   let r = fs::canonicalize("index.md");
   let s = r.unwrap_or_default();
   println!("{:?}", s);
}
