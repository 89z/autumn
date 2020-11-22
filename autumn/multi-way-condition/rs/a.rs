fn main() {
   let n = match '\x43' {
      'A' => 0x41,
      'B' | 'b' => 0x42,
      c => c as u8
   };

   println!("{}", n == 0x43);
}
