fn f(s_acc: String, s_cur: &&str) -> String {
   s_acc + s_cur
}

fn main() {
   let a = vec!["May", "June"];
   // example 1
   let s = a.iter().fold(String::new(), f);
   // example 2
   let mut s2 = String::new();
   for s_cur in a.iter() {
      s2 = f(s2, s_cur);
   }
   // print
   println!("{}", s == "MayJune" && s2 == "MayJune");
}
