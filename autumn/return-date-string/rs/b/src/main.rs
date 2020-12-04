use chrono::offset::Local;

fn main() {
   let o = Local::now();
   let s = o.to_rfc3339();
   println!("{}", s);
}
