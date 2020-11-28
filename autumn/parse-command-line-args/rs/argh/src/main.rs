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
   let pre_s = o.prefix.unwrap_or_default();
   let suf_s = o.suffix.unwrap_or_default();
   println!("{}", pre_s + &o.stem + &suf_s);
}
