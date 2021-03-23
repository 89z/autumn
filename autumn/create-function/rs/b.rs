// example 1
fn f(d: u8, e: u8) -> u8 {
   return d + e
}

// example 2
fn g(d: u8, e: u8) -> u8 {
   d + e
}

// print
fn main() {
   println!("{}", f(4, 5) == 9);
   println!("{}", g(4, 5) == 9);
}
