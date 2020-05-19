#[link(name = "ntdll")]
extern {
   fn RtlGetNtVersionNumbers(major: *mut u32, minor: *mut u32, build: *mut u32);
}
fn main() {
   let mut n1 = 0u32;
   let mut n2 = 0u32;
   let mut n3 = 0u32;
   unsafe {
      RtlGetNtVersionNumbers(&mut n1, &mut n2, &mut n3);
   }
   let n_loword = n3 & 0xFFFF;
   dbg!(n1, n2, n_loword);
}
