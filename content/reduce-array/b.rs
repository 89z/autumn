fn f(sa: String, sc: &&str) -> String {
   return sa + sc;
}

fn main() {
   let a = vec!["May", "June"];
   let mut s = String::new();
   for sc in a.iter() {
      s = f(s, sc);
   }
   println!("{}", s == "MayJune");
}
