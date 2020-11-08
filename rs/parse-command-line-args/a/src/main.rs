use pico_args::Arguments;
use std::process;

fn main() {
   let mut o = Arguments::from_env();

   let s_start: String = o.opt_value_from_str("-s").
   unwrap_or(None).unwrap_or_default();

   let s_end: String = o.opt_value_from_str("-e").
   unwrap_or(None).unwrap_or_default();

   let a = o.free().unwrap_or_default();

   if a.len() != 1 {
      println!("cat [flags] <input>
-s string
   start
-e string
   end");
      process::exit(1);
   }

   let s_in = &a[0];
   println!("{}", s_start + &s_in + &s_end);
}
