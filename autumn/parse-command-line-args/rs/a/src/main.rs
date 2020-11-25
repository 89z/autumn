use {
   pico_args::Arguments,
   std::process
};

fn main() {
   let mut o = Arguments::from_env();

   let pre_s: String = o.opt_value_from_str("-p").
   unwrap_or(None).unwrap_or_default();

   let suf_s: String = o.opt_value_from_str("-s").
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

   let stem_s = &a[0];
   println!("{}", pre_s + &stem_s + &suf_s);
}
