fn main() {
   let a = ["March", "April"];
   let e = a.iter().nth(0);
   println!("{}", e == Some(&"March"));
}
