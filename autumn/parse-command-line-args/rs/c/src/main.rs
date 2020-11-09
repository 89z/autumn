use structopt::StructOpt;

#[derive(StructOpt)]
struct Opt {
   #[structopt(short, required = false)]
   p: String,
   #[structopt(short, required = false)]
   s: String,
   #[structopt()]
   stem: String
}

fn main() {
   let o = Opt::from_args();
   println!("{}", o.p + &o.stem + &o.s);
}
