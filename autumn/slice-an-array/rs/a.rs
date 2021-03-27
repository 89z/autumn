fn main() {
   let a = [0, 10];
   let n = a.iter().nth(1);
   println!("{}", n == Some(&10));
}
