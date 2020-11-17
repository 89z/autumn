fn number_format(n: f64) -> String {
   let n2 = n.log2() / 10f64;
   let n = n / 1024f64.powi(n2 as i32);
   format!("{:.3}", n) + ["", " k", " M", " G"][n2 as usize]
}

fn main() {
   let s = number_format(1264.);
   println!("{}", s);
}
