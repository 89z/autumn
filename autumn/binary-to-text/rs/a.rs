use core::fmt::Write;

fn main() -> core::fmt::Result {
    let a = vec![22, 4, 127, 193];
    let n = a.len();
    
    let mut s = String::with_capacity(2 * n);
    for byte in a {
        write!(s, "{:02X}", byte)?;
    }
    
    dbg!(s);
    Ok(())
}
