// example 1
fn add(n: u8, n1: u8) -> u8 {
   return n + n1;
}

// example 2
fn sub(n: u8, n2: u8) -> u8 {
   n - n2
}

// print
fn main() {
   let n1 = add(8, 1);
   let n2 = sub(8, 1);
   println!("{}", n1 == 9 && n2 == 7);
}
