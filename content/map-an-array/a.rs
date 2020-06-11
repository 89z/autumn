fn main() {
   let a1 = [10, 11];
   let a2: Vec<u8> = a1.iter().map(|n| n * 2).collect();
   println!("{:?}", a2);
}
