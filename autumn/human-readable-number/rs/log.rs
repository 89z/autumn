fn number_format(n: f64) -> String {
   let n2 = n.log(1000f64);
   let n = n / 1000f64.powi(n2 as i32);
   format!("{:.3}", n) + ["", " k", " M", " G"][n2 as usize]
}

fn main() {
   let s = number_format(9012345678f64);
   println!("{}", s == "9.012 G");
}
