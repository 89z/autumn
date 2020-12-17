use {
   getopts::Options,
   std::env
};

fn main() -> Result<(), getopts::Fail> {
   let a: Vec<String> = env::args().collect();
   let mut opt_o = Options::new();
   opt_o.optopt("p", "prefix", "place before", "STRING");
   opt_o.optopt("s", "suffix", "place after", "STRING");
   let match_o = opt_o.parse(&a[1..])?;

   if match_o.free.len() != 1 {
      println!("{}", opt_o.usage("add [options] <stem>"));
      return Ok(())
   }

   let pre_s = match_o.opt_str("p").unwrap_or_default();
   let suf_s = match_o.opt_str("s").unwrap_or_default();
   let stem_s = &match_o.free[0];
   println!("{}", pre_s + stem_s + &suf_s);
   Ok(())
}
