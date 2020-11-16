fn format_size(n: u64) -> String {
    if n < 1024 {
        return format!("{}B", n);
    }
    // i is the largest value such that 1024 ^ i < n
    // To find i we compute the smallest b such that n <= 1024 ^ b and subtract 1 from it
    let i = (64 - n.leading_zeros() + 9) / 10 - 1;
    // Compute the size with two decimal places (rounded down) as the last two digits of q
    // This avoid float formatting reducing the binary size
    let q = 100 * n / (1024 << ((i - 1) * 10));
    format!("{}.{:02}{}B", q / 100, q % 100, b" kMGTPEZ"[i as usize] as char)
}

fn main() {
   let s = format_size(1264);
   println!("{}", s);
}
