fn number_format(d: f64) -> String {
   let (mut e, mut f) = (d, 0);
   while e >= 1e3 {
      e /= 1e3;
      f += 1;
   }
   format!("{:.3}", e) + ["", " k", " M", " G"][f]
}

fn main() {
   let s = number_format(9012345678.0);
   println!("{}", s == "9.012 G");
}
