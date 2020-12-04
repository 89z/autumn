use std::{
   io::Write,
   time::Duration,
   time::Instant
};

fn main() -> std::io::Result<()> {
   let now_o = Instant::now();
   let dur_o = Duration::from_millis(10);
   loop {
      std::thread::sleep(dur_o);
      let n = now_o.elapsed().as_secs_f32();
      print!("{:.2}\r", n);
      std::io::stdout().flush()?;
   }
}
