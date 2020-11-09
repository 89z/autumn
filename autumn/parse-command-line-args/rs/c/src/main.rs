use structopt::StructOpt;

#[derive(StructOpt)]
struct Opt {
   /// prefix
   #[structopt(short)]
   p: Option<String>,
   /// suffix
   #[structopt(short)]
   s: Option<String>,
   #[structopt()]
   stem: String
}

fn main() {
   let o = Opt::from_args();
   let s_pre = o.p.unwrap_or_default();
   let s_suf = o.s.unwrap_or_default();
   println!("{}", s_pre + &o.stem + &s_suf);
}
