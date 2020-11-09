// (Full example with detailed comments in examples/01a_quick_example.rs)
//
// This example demonstrates clap's "usage strings" method of creating arguments
// which is less verbose
use clap::{Arg, App};

fn main() {
    let matches = App::new("myapp")
        .version("1.0")
        .author("Kevin K. <kbknapp@gmail.com>")
        .about("Does awesome things")
        .arg("-c, --config=[FILE] 'Sets a custom config file'")
        .arg("<INPUT>              'Sets the input file to use'")
        .arg("-v...                'Sets the level of verbosity'")
        .subcommand(App::new("test")
            .about("controls testing features")
            .version("1.3")
            .author("Someone E. <someone_else@other.com>")
            .arg("-d, --debug 'Print debug information'"))
        .get_matches();
}
