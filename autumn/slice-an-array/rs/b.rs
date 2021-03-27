fn main() {
   let a = [0, 10, 20, 30, 40];
   // example 1
   let n = a.get(2);
   println!("{}", n == Some(&20));
   // example 2
   let b = a.get(2 ..);
   println!("{}", b == Some(&[20, 30, 40]));
   // example 3
   let b = a.get(2 .. 4);
   println!("{}", b == Some(&[20, 30]));
   // example 4
   let b = a.get(2 ..= 3);
   println!("{}", b == Some(&[20, 30]));
}
