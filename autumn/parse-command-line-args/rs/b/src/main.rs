use argh::FromArgs;

/// Add
#[derive(FromArgs)]
struct Args {
   /// before stem
   #[argh(option, short = 'p')]
   prefix: Option<String>,
   /// after stem
   #[argh(option, short = 's')]
   suffix: Option<String>,
   #[argh(positional)]
   stem: String
}

fn main() {
   let o: Args = argh::from_env();
   let s_pre = o.prefix.unwrap_or_default();
   let s_suf = o.suffix.unwrap_or_default();
   println!("{}", s_pre + &o.stem + &s_suf);
}
