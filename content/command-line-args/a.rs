use std::env::args;

fn main() {
   let a: Vec<String> = args().collect();
   let s = &a[1];
   println!("{}", s == "May");
}
