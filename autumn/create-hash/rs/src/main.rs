use hex_literal::hex;
use md5::{Md5, Digest};

fn main() {
   let mut hasher = Md5::new();
   hasher.update(b"hello world");
   let result = hasher.finalize();
   assert_eq!(result[..], hex!("5eb63bbbe01eeed093cb22bb8f5acdc3"));
}
