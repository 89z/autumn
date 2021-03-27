use chrono::offset::Utc;

fn main() {
   let d = Utc::today();
   println!("{}", d);
}
