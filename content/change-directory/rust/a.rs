use std::env;

fn main() {
   let r = env::set_current_dir("..");
   println!("{:?}", r);
}
