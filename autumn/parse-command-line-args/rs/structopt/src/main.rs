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
   let pre_s = o.p.unwrap_or_default();
   let suf_s = o.s.unwrap_or_default();
   println!("{}", pre_s + &o.stem + &suf_s);
}
