fn main() {
   let c = match 0x63 - 0x20u8 {
      0x41 => 'A',
      0x42 | 0x62 => 'B',
      n => n as char
   };
   println!("{}", c == 'C');
}
