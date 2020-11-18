fn number_format(n: f64) -> String {
   let mut n2 = n;
   let mut n3 = 0;
   while n2 >= 1e3 {
      n2 /= 1e3;
      n3 += 1;
   }
   format!("{:.3}", n2) + ["", " k", " M", " G"][n3]
}

fn main() {
   let s = number_format(9012345678e0);
   println!("{}", s == "9.012 G");
}
