use data_encoding::BASE64;

fn main() {
   let s = [10, 11, 12, 13];
   let t = BASE64.encode(&s);
   println!("{}", t == "CgsMDQ==");
}
