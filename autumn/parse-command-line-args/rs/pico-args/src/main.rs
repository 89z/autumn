use {
   pico_args::Arguments,
   std::process
};

fn main() {
   let mut arg = Arguments::from_env();
   let pre: String = arg.value_from_str("-p").unwrap_or_default();
   let suf: String = arg.value_from_str("-s").unwrap_or_default();
   let a = arg.free().unwrap_or_default();

   if a.len() != 1 {
      println!("add [flags] <stem>
-p string
   prefix
-s string
   suffix");
      process::exit(1);
   }

   let stem = &a[0];
   println!("{}", pre + &stem + &suf);
}
