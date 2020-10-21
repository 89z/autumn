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
    // Arguments can be parsed in any order.
    let args = Args {
        // You can use a slice for multiple commands
        help: args.contains(["-h", "--help"]),
        // or just a string for a single one.
        version: args.contains("-V"),
        // Parses an optional value that implements `FromStr`.
        number: args.opt_value_from_str("--number")?.unwrap_or(5),
        // Parses an optional value that implements `FromStr`.
        opt_number: args.opt_value_from_str("--opt-number")?,
        // Parses an optional value using a specified function.
        width: args.opt_value_from_fn("--width", parse_width)?.unwrap_or(10),
        // Will return all free arguments or an error if any flags are left.
        free: args.free()?,
    };
    Ok(())
}
