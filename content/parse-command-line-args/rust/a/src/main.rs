struct Args {
   help: bool,
   version: bool,
   number: u32,
   opt_number: Option<u32>,
   width: u32,
   free: Vec<String>,
}

fn parse_width(s: &str) -> Result<u32, String> {
   s.parse().map_err(|_| "not a number".to_string())
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
   let mut args = pico_args::Arguments::from_env();
   let args = Args {
      help: args.contains(["-h", "--help"]),
      version: args.contains("-V"),
      number: args.opt_value_from_str("--number")?.unwrap_or(5),
      opt_number: args.opt_value_from_str("--opt-number")?,
      width: args.opt_value_from_fn("--width", parse_width)?.unwrap_or(10),
      free: args.free()?,
   };
   Ok(())
}
