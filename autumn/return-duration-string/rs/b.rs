use std::{
   error::Error,
   io::Write,
   time::Duration,
   time::SystemTime
};

fn main() -> Result<(), Box<dyn Error>> {
   let then = SystemTime::now();
   let dur = Duration::from_millis(10);
   loop {
      std::thread::sleep(dur);
      let now = then.elapsed()?;
      let now = now.as_secs_f32();
      print!("\r{:.2}", now);
      std::io::stdout().flush()?;
   }
}
