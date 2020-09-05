#![allow(unused)]
use std::fs;

fn main() -> std::io::Result<()> {
    fs::copy("foo.txt", "bar.txt")?;  // Copy foo.txt to bar.txt
    Ok(())
}
