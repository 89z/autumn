use std::fs;

fn main() {
   let e = fs::read_to_string("index.md");
   let s = e.unwrap_or_default();
   print!("{}", s);
}
