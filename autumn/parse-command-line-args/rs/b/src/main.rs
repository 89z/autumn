use argh::FromArgs;

/// concatenate
#[derive(FromArgs)]
struct Args {
   /// prefix
   #[argh(option)]
   start: Option<String>,
   /// suffix
   #[argh(option)]
   end: Option<String>,
   #[argh(positional)]
   input: String
}

fn main() {
   let o: Args = argh::from_env();
   let s_start = o.start.unwrap_or_default();
   let s_end = o.end.unwrap_or_default();
   println!("{}", s_start + &o.input + &s_end);
}
