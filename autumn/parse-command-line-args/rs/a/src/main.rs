use {pico_args::Arguments, std::process};

fn main() {
   let mut o = Arguments::from_env();

   let s_pre: String = o.opt_value_from_str("-p").
   unwrap_or(None).unwrap_or_default();

   let s_suf: String = o.opt_value_from_str("-s").
   unwrap_or(None).unwrap_or_default();

   let a = o.free().unwrap_or_default();

   if a.len() != 1 {
      println!("add [flags] <stem>
-p string
   prefix
-s string
   suffix");
      process::exit(1);
   }

   let s_stem = &a[0];
   println!("{}", s_pre + &s_stem + &s_suf);
}
