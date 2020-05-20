use std::env;
fn main() {
   match env::var("BROWSER") {
      Err(e) => println!("{}", e),
      Ok(s1) => println!("{}", s1)
   }
}
