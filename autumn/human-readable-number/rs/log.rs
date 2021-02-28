fn number_format(f: f64) -> String {
   let g = f.log(1e3);
   let h = f / f64::powi(1e3, g as i32);
   format!("{:.3}", h) + ["", " k", " M", " G"][g as usize]
}

fn main() {
   let s = number_format(9012345678.0);
   println!("{}", s == "9.012 G");
}
