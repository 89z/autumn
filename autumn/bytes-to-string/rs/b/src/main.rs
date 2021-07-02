use data_encoding::HEXUPPER;

fn main() {
   let s = [10, 11, 12, 13];
   let t = HEXUPPER.encode(&s);
   println!("{}", t == "0A0B0C0D");
}
