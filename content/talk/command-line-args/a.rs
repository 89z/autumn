use std::env;
fn main() {
   let a: Vec<String> = env::args().collect();
   let s = &a[1];
   println!("{}", s == r#"a "b"#);
}
