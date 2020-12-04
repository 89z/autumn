use std::{
   error::Error,
   io::Write,
   time::Duration,
   time::SystemTime
};

fn main() -> Result<(), Box<dyn Error>> {
   let old_o = SystemTime::now();
   let dur_o = Duration::from_millis(10);
   loop {
      std::thread::sleep(dur_o);
      let new_o = old_o.elapsed()?;
      let n = new_o.as_secs_f32();
      print!("{:.2}\r", n);
      std::io::stdout().flush()?;
   }
}
