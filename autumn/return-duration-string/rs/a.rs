use std::{
   io::Write,
   time::Duration,
   time::Instant
};

fn main() -> std::io::Result<()> {
   let then = Instant::now();
   let dur = Duration::from_millis(10);
   loop {
      std::thread::sleep(dur);
      let now = then.elapsed().as_secs_f32();
      print!("{:.2}\r", now);
      std::io::stdout().flush()?;
   }
}
