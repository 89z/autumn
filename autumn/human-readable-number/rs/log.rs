fn number_format(n: f64) -> String {
   let n2 = n.log(1e3);
   let n = n / 1e3f64.powi(n2 as i32);
   format!("{:.3}", n) + ["", " k", " M", " G"][n2 as usize]
}

fn main() {
   let s = number_format(9012345678e0);
   println!("{}", s == "9.012 G");
}
