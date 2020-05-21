use regex::Regex;
fn main() {
   let s1 = "Wednesday";
   match Regex::new(r"e(..)") {
      Err(e) => {
         dbg!(e);
      },
      Ok(o1) => for a1 in o1.captures_iter(s1) {
         dbg!(a1);
      }
   }
}
