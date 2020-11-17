fn number_format(mut n: f64) -> String {
   let mut n2 = 0;
   while n >= 1e3 {
      n /= 1e3;
      n2 += 1;
   }
   format!("{:.3}", n) + ["", " k", " M", " G"][n2]
}

fn main() {
   let s = number_format(9012345678e0);
   println!("{}", s == "9.012 G");
}
