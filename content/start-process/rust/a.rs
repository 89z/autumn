use std::process::Command;

fn main() {
   Command::new("cc").arg("--print-search-dirs");
}
