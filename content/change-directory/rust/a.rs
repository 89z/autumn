use std::env;

fn main() {
   let e = env::set_current_dir("..");
   println!("{:?}", e);
}
