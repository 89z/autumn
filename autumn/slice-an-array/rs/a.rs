fn main() {
   let a = [9, 8];
   let n = a.iter().nth(1);
   println!("{}", n == Some(&8));
}
