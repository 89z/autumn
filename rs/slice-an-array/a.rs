fn main() {
   let a = ["May", "June"];
   let e = a.iter().nth(0);
   println!("{}", e == Some(&"May"));
}
