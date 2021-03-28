use chrono::NaiveDateTime;

fn main() {
   let n = 0x5555_5555;
   let d = NaiveDateTime::from_timestamp(n, 0);
   println!("{}", d);
}
