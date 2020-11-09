use clap::App;

fn main() {
   let o = App::new("cat").
   arg("<input> 'input file'").
   arg("-s [string] 'start'").
   arg("-e [string] 'end'").
   get_matches();

   let s_start = o.value_of_lossy("s").unwrap_or_default();
   let s_end = o.value_of_lossy("e").unwrap_or_default();
   let s_input = o.value_of_lossy("input").unwrap_or_default();
   println!("{}", s_start + s_input + s_end);
}
