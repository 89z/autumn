use chrono::offset::Local;

fn main() {
   let d = Local::now();
   let s = d.to_rfc2822();
   println!("{}", s);
}
