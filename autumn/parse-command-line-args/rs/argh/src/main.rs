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
   let arg: Args = argh::from_env();
   let pre = arg.prefix.unwrap_or_default();
   let suf = arg.suffix.unwrap_or_default();
   println!("{}", pre + &arg.stem + &suf);
}
