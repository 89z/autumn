use std::{
   io::Write,
   time::Duration,
   time::Instant
};

fn main() -> std::io::Result<()> {
   let now = Instant::now();
   let dur = Duration::from_millis(10);
   loop {
      std::thread::sleep(dur);
      let n = now.elapsed().as_secs_f32();
      print!("{:.2}\r", n);
      std::io::stdout().flush()?;
   }
}
