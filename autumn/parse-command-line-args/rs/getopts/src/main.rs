use {
   getopts::Options,
   std::env
};

fn main() -> Result<(), getopts::Fail> {
   let a: Vec<String> = env::args().collect();
   let mut opt = Options::new();
   opt.optopt("p", "prefix", "place before", "STRING");
   opt.optopt("s", "suffix", "place after", "STRING");
   let parse = opt.parse(&a[1..])?;

   if parse.free.len() != 1 {
      println!("{}", opt.usage("add [options] <stem>"));
      return Ok(())
   }

   let pre = parse.opt_str("p").unwrap_or_default();
   let suf = parse.opt_str("s").unwrap_or_default();
   let stem = &parse.free[0];
   println!("{}", pre + stem + &suf);
   Ok(())
}
