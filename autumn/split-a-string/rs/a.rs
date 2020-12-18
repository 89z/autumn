fn split(s: &str, sep: char) -> Vec<&str> {
   s.split(sep).collect()
}

fn main() {
   let s = "May,June";
   let a = split(s, ',');
   println!("{:?}", a);
}
